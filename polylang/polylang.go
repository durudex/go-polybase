/*
 * Copyright © 2022 V1def
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polylang

import "github.com/alecthomas/participle/v2/lexer"

type Polylang struct {
	Pos lexer.Position

	Collection *Collection
}

type Collection struct {
	Pos lexer.Position

	Name  string  `parser:"'collection' @Ident"`
	Items []*Item `parser:"'{' @@* '}'"`
}

type Item struct {
	Pos lexer.Position

	Field    *Field    `parser:"@@"`
	Function *Function `parser:"| @@"`
	Index    *Index    `parser:"| @@"`
}

type Field struct {
	Pos lexer.Position

	Name     string `parser:"@Ident"`
	Optional bool   `parser:"@'?'?"`
	Type     string `parser:"':' @Ident ';'"`
}

type Function struct {
	Pos lexer.Position

	Name       string       `parser:"@Ident '('"`
	Parameters []*Parameter `parser:"( @@ ( ',' @@ )* )? ')'"`
	ReturnType string       `parser:"( ':' @Ident )?"`
	Statements []*Statement `parser:"'{' ( @@* )? '}'"`
}

type Parameter struct {
	Pos lexer.Position

	Name     string `parser:"@Ident"`
	Optional bool   `parser:"@'?'?"`
	Type     string `parser:"':' @Ident"`
}

type Index struct {
	Pos lexer.Position

	Unique bool          `parser:"'@' ( @'unique' | 'index' )"`
	Fields []*IndexField `parser:"'(' @@ ( ',' @@ )* ')' ';'"`
}

type IndexField struct {
	Pos lexer.Position

	Name  string `parser:"( '[' )? ( @Ident )"`
	Order string `parser:"( ',' @Ident ']' )?"`
}

type Statement struct {
	Pos lexer.Position

	Expression *Expression `parser:"@@ ';'"`
	If         *If         `parser:"| @@"`
}

type Expression struct {
	Pos lexer.Position

	Left     string `parser:"@Ident"`
	Operator string `parser:"@( '=' '=' | '!' '=' | '=' )"`
	Right    string `parser:"@Ident"`
}

type If struct {
	Pos lexer.Position

	Expression *Expression  `parser:"'if' '(' @@ ')'"`
	Statements []*Statement `parser:"'{' @@* '}'"`
	Else       []*Statement `parser:"( 'else' '{' @@* '}' )?"`
}
