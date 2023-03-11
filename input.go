/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polybase

import "reflect"

var allowedKindTypes = map[reflect.Kind]bool{
	reflect.String: true, reflect.Int: true, reflect.Int8: true,
	reflect.Int16: true, reflect.Int32: true, reflect.Int64: true,
	reflect.Uint: true, reflect.Uint8: true, reflect.Uint16: true,
	reflect.Uint32: true, reflect.Uint64: true, reflect.Bool: true,
}

func ParseInput(args []any) []any {
	var res []any

	for _, arg := range args {
		val := reflect.ValueOf(arg)

		switch val.Kind() {
		case reflect.Array, reflect.Slice:
			switch val.Type().Elem().Kind() {
			case reflect.Interface:
				res = append(res, arg.([]any)...)
			default:
				if !allowedKindTypes[val.Type().Elem().Kind()] {
					continue
				}

				piv := parseIterableValue(val)
				res = append(res, piv...)
			}
		case reflect.Ptr:
			switch val.Elem().Kind() {
			case reflect.Struct:
				psv := parseStructValue(val.Elem())
				res = append(res, psv...)
			default:
				continue
			}
		case reflect.Struct:
			psv := parseStructValue(val)
			res = append(res, psv...)
		default:
			if !allowedKindTypes[val.Kind()] {
				continue
			}

			res = append(res, arg)
		}
	}

	return res
}

func parseIterableValue(v reflect.Value) []any {
	n := v.Len()
	res := make([]any, 0, n)

	for i := 0; i < n; i++ {
		res = append(res, v.Index(i).Interface())
	}

	return res
}

func parseStructValue(v reflect.Value) []any {
	n := v.NumField()
	res := make([]any, 0, n)

	for i := 0; i < n; i++ {
		field := v.Field(i)

		if !allowedKindTypes[field.Kind()] {
			continue
		}

		res = append(res, field.Interface())
	}

	return res
}
