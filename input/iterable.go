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

	if e.Kind() == reflect.Interface {
		return v.Interface().([]any)
	} else if !AllowedKindTypes[e.Kind()] {
		panic("error: unsupported type")
	}

	return []any{v.Interface()}
}
