/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package input

import "reflect"

func Parse(args []any) []any {
	var res []any

	for _, arg := range args {
		val := reflect.ValueOf(arg)

		switch val.Kind() {
		case reflect.Array, reflect.Slice:
			pv := parseIterableValue(&val)

			if res == nil {
				res = pv
				continue
			}

			res = append(res, pv...)
		case reflect.Struct:
			pf := parseForeignValue(&val)
			if pf != nil {
				res = append(res, pf)
				continue
			}

			pv := parseStructValue(&val)

			if res == nil {
				res = pv
				continue
			}

			res = append(res, pv...)
		case reflect.Ptr:
			pv := parsePointerValue(&val)

			if res == nil {
				res = pv
				continue
			}

			res = append(res, pv...)
		default:
			if !AllowedKindTypes[val.Kind()] {
				panic("error: unsupported type")
			}

			res = append(res, arg)
		}
	}

	return res
}
