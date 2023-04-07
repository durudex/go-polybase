/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the LICENSE
 * file in the root directory of this source tree.
 */

package polybase

import (
	"context"
	"log"
)

type RecoverHandler = func(ctx context.Context, v any)

func DefaultRecover(ctx context.Context, v any) { log.Fatalf("FATAL: %v", v) }

func recoverFunc(ctx context.Context, rf RecoverHandler) {
	if v := recover(); v != nil {
		rf(ctx, v)
	}
}
