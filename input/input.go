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
		v := reflect.ValueOf(arg)

		switch v.Kind() {
		case reflect.Array, reflect.Slice:
			pv := parseArrayValue(&v)

			if res == nil {
				res = pv
				continue
			}

			res = append(res, pv...)
		case reflect.Map:
			pv := parseMapValue(&v)
			res = append(res, pv)
		case reflect.Struct:
			pf := parseForeignValue(&v)
			if pf != nil {
				res = append(res, pf)
				continue
			}

			pv := parseStructValue(&v)

			if res == nil {
				res = pv
				continue
			}

			res = append(res, pv...)
		case reflect.Ptr:
			pv := parsePointerValue(&v)

			if res == nil {
				res = pv
				continue
			}

			res = append(res, pv...)
		default:
			if !AllowedKindType[v.Kind()] {
				panic("error: unsupported type")
			}

			res = append(res, arg)
		}
	}

	return res
}
