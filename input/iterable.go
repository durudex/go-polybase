/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package input

import "reflect"

func ParseIterable(arg any) []any {
	val := reflect.ValueOf(arg)
	return parseIterableValue(&val)
}

func parseIterableValue(v *reflect.Value) []any {
	e := v.Type().Elem()

	switch e.Kind() {
	case reflect.Interface:
		return v.Interface().([]any)
	case reflect.Struct:
		n := v.Len()
		res := make([]any, n)

		for i := 0; i < n; i++ {
			field := v.Index(i)

			pv := parseForeignValue(&field)
			if pv == nil {
				panic("error: unsupported nested struct")
			}

			res[i] = pv
		}

		return []any{res}
	}

	if !AllowedKindTypes[e.Kind()] {
		panic("error: unsupported type")
	}

	return []any{v.Interface()}
}
