/*
 * Copyright © 2022 V1def
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polybase

func buildParam(params []string) string {
	n := len(params)

	if n != 0 {
		res := "?"

		for _, par := range params[:n-1] {
			res += par + "&"
		}

		res += params[n-1]

		return res
	}

	return ""
}
