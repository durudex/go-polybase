/*
 * Copyright © 2022 V1def
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polybase

import "reflect"

func ParseInput(v any) []any {
	e := reflect.ValueOf(v).Elem()

	if e.Type().Kind() != reflect.Struct {
		return nil
	}

	res := make([]any, e.NumField())

	for i := 0; i < e.NumField(); i++ {
		res[i] = e.Field(i).Interface()

	}

	return res
}
