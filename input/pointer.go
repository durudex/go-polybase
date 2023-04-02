/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package input

import "reflect"

func ParsePointer(arg any) []any {
	val := reflect.ValueOf(arg)
	return parsePointerValue(&val)
}

func parsePointerValue(v *reflect.Value) []any {
	e := v.Elem()

	if v.IsNil() {
		panic("error: unsupported nil value")
	} else if e.Kind() == reflect.Struct {
		pf := parseForeignValue(v)
		if pf != nil {
			return []any{pf}
		}

		return parseStructValue(&e)
	} else if !AllowedKindType[e.Kind()] {
		panic("error: unsupported type")
	}

	return []any{e.Interface()}
}
