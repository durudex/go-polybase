/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polybase

type Code string

const (
	CodeInvalidArgument    Code = "invalid-argument"
	CodeFailedPrecondition Code = "failed-precondition"
	CodeOutOfRange         Code = "out-of-range"
	CodeUnauthenticated    Code = "unauthenticated"
	CodePermissionDenied   Code = "permission-denied"
	CodeNotFound           Code = "not-found"
	CodeAborted            Code = "aborted"
	CodeAlreadyExists      Code = "already-exists"
	CodeResourceExhausted  Code = "resource-exhausted"
	CodeCancelled          Code = "cancelled"
	CodeUnavailable        Code = "unavailable"
	CodeInternal           Code = "internal"
	CodeDeadlineExceeded   Code = "deadline-exceeded"
)

type Error struct {
	Reason  string `json:"reason"`
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string { return e.Message }
