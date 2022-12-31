/*
 * Copyright © 2022 Durudex
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

	"github.com/durudex/go-polybase"
	"github.com/durudex/go-polybase/codegen/template"
	"github.com/durudex/go-polybase/polylang"

	"github.com/alecthomas/participle/v2"
	"github.com/iancoleman/strcase"
)

const GenesisCollectionID = "Collection"

type GenesisCollection struct {
	Code string `json:"code"`
}

type ParsedCollection struct {
	ID        string
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
	names := make([]string, len(c.config.Collections))

	if err := c.checkDir(); err != nil {
		return err
	}

	for i, id := range c.config.Collections {
		ast, err := c.astCollection(context.Background(), id)
		if err != nil {
			return err
		}

		parsed := c.parseAst(ast)
		parsed.ID = id

		names[i] = parsed.Name

		if err := c.generateFile(parsed); err != nil {
			return err
		}
	}

	if err := c.generateClient(names); err != nil {
		return err
	}

	return c.fmt()
}

func (c *codegen) generateClient(names []string) error {
	f, err := os.Create(c.config.Directory + "/client.go")
	if err != nil {
		return err
	}
	defer f.Close()

	template.WriteHeader(f, c.config.Package)
	template.WriteClient(f, names)

	return nil
}

func (c *codegen) generateFile(coll *ParsedCollection) error {
	f, err := os.Create(c.config.Directory + "/" + strcase.ToSnake(coll.Name) + ".go")
	if err != nil {
		return err
	}
	defer f.Close()

	template.WriteHeader(f, c.config.Package)
	template.WriteImport(f)
	template.WriteModel(f, coll.Name, coll.Fields)
	template.WriteCollection(f, coll.ID, coll.Name, coll.Functions)

	for _, fc := range coll.Functions {
		template.WriteInput(f, coll.Name, fc.Name, fc.Parameters)
		template.WriteFunction(f, coll.Name, fc)
	}

	return nil
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
