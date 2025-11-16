// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

func ReturnPointer[S any](s S) *S {
	return &s
}
