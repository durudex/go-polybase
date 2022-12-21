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
	"os/exec"
	"path/filepath"

	"github.com/v1def/go-polybase"
	"github.com/v1def/go-polybase/codegen/template"
	"github.com/v1def/go-polybase/polylang"

	"github.com/alecthomas/participle/v2"
	"github.com/iancoleman/strcase"
)

const GenesisCollectionID = "Collection"

type GenesisCollection struct {
	Code string `json:"code"`
}

type ParsedCollection struct {
	Name      string
	Fields    []*polylang.Field
	Functions []*polylang.Function
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

		if err := c.generateFile(c.parseAst(ast)); err != nil {
			return err
		}
	}

	return nil
}

func (c *codegen) generateFile(coll *ParsedCollection) error {
	if err := c.checkDir(); err != nil {
		return err
	}

	path := c.config.Directory + "/" + strcase.ToSnake(coll.Name) + ".go"

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	template.WriteHeader(f, c.config.Package)
	template.WriteModel(f, coll.Name, coll.Fields)
	template.WriteFunction(f, coll.Name, coll.Functions)

	return c.fmt()
}

func (c *codegen) fmt() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	return exec.Command("go", "fmt", filepath.Join(dir, c.config.Directory)).Run()
}

func (c *codegen) checkDir() error {
	if err := os.MkdirAll(c.config.Directory, 0755); err != nil {
		if os.IsExist(err) {
			// TODO: add delete old files by option
			return nil
		}

		return err
	}

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
		Name:      ast.Name,
		Fields:    make([]*polylang.Field, 0),
		Functions: make([]*polylang.Function, 0),
	}

	for _, item := range ast.Items {
		switch {
		case item.Field != nil:
			collection.Fields = append(collection.Fields, item.Field)
		case item.Function != nil:
			collection.Functions = append(collection.Functions, &polylang.Function{
				Name:       item.Function.Name,
				Parameters: item.Function.Parameters,
			})
		}
	}

	return collection
}
