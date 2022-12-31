/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polylang

import "github.com/alecthomas/participle/v2/lexer"

var PolylangLexer = lexer.MustSimple([]lexer.SimpleRule{
	{Name: "Ident", Pattern: `[a-zA-Z_.][a-zA-Z0-9_.]*`},
	{Name: "String", Pattern: `'[^']*'`},
	{Name: "Punct", Pattern: `\[|]|[,:;={}()?!@]`},
	{Name: "whitespace", Pattern: `\s+`},
})
