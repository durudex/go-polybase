/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package input

import "reflect"

func ParseStruct(arg any) []any {
	val := reflect.ValueOf(arg)
	return parseStructValue(&val)
}

func parseStructValue(v *reflect.Value) []any {
	n := v.NumField()
	res := make([]any, 0, n)

	for i := 0; i < n; i++ {
		field := v.Field(i)

		switch field.Kind() {
		case reflect.Pointer:
			pv := parsePointerValue(&field)
			res = append(res, pv...)
		case reflect.Struct:
			pv := parseForeignValue(&field)
			if pv == nil {
				panic("error: unsupported nested struct")
			}

			res = append(res, pv)
		default:
			if !AllowedKindTypes[field.Kind()] {
				panic("error: unsupported type")
			}

			res = append(res, field.Interface())
		}
	}

	return res
}
