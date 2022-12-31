/*
 * Copyright © 2022 Durudex
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
	Order Order  `parser:"( ',' @@ ']' )?"`
}

type Statement struct {
	CompoundStatement
	SimpleStatement
}

type CompoundStatement struct {
	If *If `parser:"@@"`
}

type SimpleStatement struct {
	Small *SmallStatement `parser:"| @@ ';'"`
}

type SmallStatement struct {
	Throw      *Expression `parser:"'throw' @@"`
	Expression *Expression `parser:"| @@"`
}

type Expression struct {
	Pos lexer.Position

	Left       string      `parser:"@( Ident | String )"`
	Operator   Operator    `parser:"( @@ )?"`
	Expression *Expression `parser:"( '(' @@ ')' )?"`
	Right      string      `parser:"( @( Ident | String ) )?"`
}

type If struct {
	Pos lexer.Position

	Condition  *Expression  `parser:"'if' '(' @@ ')'"`
	Statements []*Statement `parser:"'{' @@* '}'"`
	Else       []*Statement `parser:"( 'else' '{' @@* '}' )?"`
}
