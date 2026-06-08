// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

func ReturnPointer[S any](s S) *S {
	return &s
}
