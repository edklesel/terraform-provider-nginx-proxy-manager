// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package go_npm_proxy_manager

import (
	"testing"
)

func TestRedirectionHostCRUD(t *testing.T) {

	var client = newTestClient(t)

	testCRUD(t, client.RedirectionHosts)
}
