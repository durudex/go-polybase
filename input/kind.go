/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the LICENSE
 * file in the root directory of this source tree.
 */

package input

import "reflect"

var (
	AllowedKindType = [...]bool{
		reflect.String: true, reflect.Int: true, reflect.Int8: true,
		reflect.Int16: true, reflect.Int32: true, reflect.Int64: true,
		reflect.Uint: true, reflect.Uint8: true, reflect.Uint16: true,
		reflect.Uint32: true, reflect.Uint64: true, reflect.Bool: true,
	}

	AllowedMapKeyKind = [...]bool{
		reflect.String: true, reflect.Int: true, reflect.Int8: true,
		reflect.Int16: true, reflect.Int32: true, reflect.Int64: true,
		reflect.Uint: true, reflect.Uint8: true, reflect.Uint16: true,
		reflect.Uint32: true, reflect.Uint64: true,
	}
)
