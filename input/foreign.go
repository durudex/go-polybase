/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the LICENSE
 * file in the root directory of this source tree.
 */

package input

import "reflect"

const customForeignMethod = "Reference"

var (
	foreignStructType  = reflect.TypeOf(Foreign{})
	foreignPointerType = reflect.TypeOf((*Foreign)(nil))
	customForeignType  = reflect.TypeOf((*CustomForeign)(nil))
)

type Foreign struct {
	CollectionID string `json:"collectionId"`
	ID           string `json:"id"`
}

type CustomForeign interface {
	Reference() *Foreign
}

func ParseForeign(arg any) any {
	val := reflect.ValueOf(arg)
	return parseForeignValue(&val)
}

func parseForeignValue(v *reflect.Value) any {
	t := v.Type()

	if t == foreignStructType || t == foreignPointerType {
		return v.Interface()
	} else if isCustomForeign(t) {
		return callForeignMethod(v).Interface()
	}

	return nil
}

func callForeignMethod(v *reflect.Value) reflect.Value {
	res := v.MethodByName(customForeignMethod).
		Call([]reflect.Value{})

	return res[0]
}

func isCustomForeign(t reflect.Type) bool {
	return t.Implements(customForeignType.Elem())
}
