/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the LICENSE
 * file in the root directory of this source tree.
 */

package input

import "reflect"

func ParseMap(arg any) any {
	val := reflect.ValueOf(arg)
	return parseMapValue(&val)
}

func parseMapValue(v *reflect.Value) any {
	t := v.Type()

	if !AllowedMapKeyKind[t.Key().Kind()] {
		panic("error: unsupported key type")
	}

	e := t.Elem()
	if e == foreignStructType || e == foreignPointerType {
		return v.Interface()
	} else if isCustomForeign(e) {
		iter := v.MapRange()
		mt := reflect.MapOf(t.Key(), foreignPointerType)
		res := reflect.MakeMapWithSize(mt, v.Len())

		for iter.Next() {
			rv := iter.Value()
			pv := callForeignMethod(&rv)

			res.SetMapIndex(iter.Key(), pv)
		}

		return res.Interface()
	} else {
		if !AllowedKindType[e.Kind()] {
			panic("error: unsupported value type")
		}

		return v.Interface()
	}
}
