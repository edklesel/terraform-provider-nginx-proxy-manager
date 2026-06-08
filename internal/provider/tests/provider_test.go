// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tests

import (
	"fmt"

	provider "github.com/edklesel/terraform-provider-nginx-proxy-manager/internal/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const (
	providerHost     = "localhost:81"
	providerUsername = "admin@example.com"
	providerPassword = "changeme"
)

var (
	providerConfig = fmt.Sprintf(`
provider "nginxproxymanager" {
  host     = "%s"
  username = "%s"
  password = "%s"
  secure   = false
}
`, providerHost, providerUsername, providerPassword)
)

var (
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"nginxproxymanager": providerserver.NewProtocol6WithError(provider.New("test")()),
	}
)
