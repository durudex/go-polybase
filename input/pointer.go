/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the LICENSE
 * file in the root directory of this source tree.
 */

package input

import "reflect"

func ParsePointer(arg any) []any {
	val := reflect.ValueOf(arg)
	return parsePointerValue(&val)
}

func parsePointerValue(v *reflect.Value) []any {
	e := v.Elem()

	pv := parseSimplePointerValue(v)
	if pv == nil {
		return parseStructValue(&e)
	}

	return []any{pv}
}

func parseSimplePointerValue(v *reflect.Value) any {
	e := v.Elem()

	if v.IsNil() {
		panic("error: unsupported pointer nil value")
	} else if e.Kind() == reflect.Struct {
		pv := parseForeignValue(v)
		if pv == nil {
			return nil
		}

		return pv
	} else if !AllowedKindType[e.Kind()] {
		panic("error: unsupported pointer type")
	}

	return v.Interface()
}
