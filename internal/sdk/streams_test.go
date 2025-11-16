// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package go_npm_proxy_manager

import (
	"testing"
)

func TestStreamCRUD(t *testing.T) {

	var client = newTestClient(t)

	testCRUD(t, client.Streams)
}
