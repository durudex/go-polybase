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

	Collections *Collection
}

type Collection struct {
	Pos lexer.Position

	Name  string  `parser:"'collection' @Ident"`
	Items []*Item `parser:"'{' @@* '}'"`
}

type Item struct {
	Pos lexer.Position

	Field *Field `parser:"@@ ';'"`
}

type Field struct {
	Pos lexer.Position

	Name     string `parser:"@Ident"`
	Required bool   `parser:"@'?'?"`
	Type     string `parser:"':' @Ident"`
}
