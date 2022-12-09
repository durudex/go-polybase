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

	Field    *Field    `parser:"@@ ';'"`
	Function *Function `parser:"@@"`
}

type Field struct {
	Pos lexer.Position

	Name     string `parser:"@Ident"`
	Optional bool   `parser:"@'?'?"`
	Type     string `parser:"':' @Ident"`
}

type Function struct {
	Pos lexer.Position

	Name           string   `parser:"@Ident '('"`
	Parameters     []*Field `parser:"( @@ ( ',' @@ )* )? ')'"`
	ReturnType     string   `parser:"( ':' @Ident )?"`
	StatementsCode string   `parser:"'{' ( @Ident )? '}'"`
}
