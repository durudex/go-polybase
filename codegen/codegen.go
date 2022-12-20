/*
 * Copyright © 2022 V1def
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package codegen

import (
	"context"
	"os"
	"strings"

	"github.com/v1def/go-polybase"
	"github.com/v1def/go-polybase/codegen/template"
	"github.com/v1def/go-polybase/polylang"

	"github.com/alecthomas/participle/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const GenesisCollectionID = "Collection"

type GenesisCollection struct {
	Code string `json:"code"`
}

type ParsedCollection struct {
	Name   string
	Fields []*polylang.Field
}

type Codegen interface {
	Generate() error
}

type codegen struct {
	config *Config
	coll   polybase.Collection
}

func New(filename string) (Codegen, error) {
	cfg, err := NewConfig(filename)
	if err != nil {
		return nil, err
	}

	client := polybase.New(polybase.Config{
		URL: polybase.TestnetURL,
	})

	return &codegen{
		config: cfg,
		coll:   client.Collection(GenesisCollectionID),
	}, nil
}

func (c *codegen) Generate() error {
	for _, id := range c.config.Collections {
		ast, err := c.astCollection(context.Background(), id)
		if err != nil {
			return err
		}

		if err := c.generate(c.parseAst(ast)); err != nil {
			return err
		}
	}

	return nil
}

func (c *codegen) generate(coll *ParsedCollection) error {
	if err := os.MkdirAll(c.config.Directory, 0755); err != nil {
		if os.IsExist(err) {
			goto gen
		}

		return err
	}

gen:

	path := c.config.Directory + "/" + strings.ToLower(coll.Name) + ".go"

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString("package " + c.config.Package + "\n")

	template.WriteModel(f, c.title(coll.Name), coll.Fields)

	return nil
}

func (c *codegen) astCollection(ctx context.Context, id string) (*polylang.Collection, error) {
	var response polybase.SingleResponse[GenesisCollection]

	if err := c.coll.Record(id).Get(ctx, &response); err != nil {
		return nil, err
	}

	parser := participle.MustBuild[polylang.Collection](
		participle.Lexer(polylang.PolylangLexer),
	)

	return parser.ParseString("", response.Data.Code)
}

func (c *codegen) parseAst(ast *polylang.Collection) *ParsedCollection {
	collection := &ParsedCollection{
		Name:   ast.Name,
		Fields: make([]*polylang.Field, 0),
	}

	for _, item := range ast.Items {
		if item.Field != nil {
			collection.Fields = append(collection.Fields, &polylang.Field{
				Name:     c.title(item.Field.Name),
				Optional: item.Field.Optional,
				Type:     item.Field.Type,
			})
		}
	}

	return collection
}

func (c *codegen) title(s string) string { return cases.Title(language.English).String(s) }
