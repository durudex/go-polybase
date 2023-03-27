/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package input

import "reflect"

type Foreign struct {
	CollectionID string `json:"collectionId"`
	ID           string `json:"id"`
}

type ForeignDoer interface {
	Reference() *Foreign
}

func ParseForeign(arg any) any {
	val := reflect.ValueOf(arg)
	return parseForeignValue(&val)
}

func parseForeignValue(v *reflect.Value) any {
	t := v.Type()

	if t == reflect.TypeOf(Foreign{}) || t == reflect.TypeOf((*Foreign)(nil)) {
		return v.Interface()
	}

	fd := reflect.TypeOf((*ForeignDoer)(nil)).Elem()

	if t.Implements(fd) {
		v := v.MethodByName("Reference").Call([]reflect.Value{})
		return v[0].Interface()
	}

	return nil
}
