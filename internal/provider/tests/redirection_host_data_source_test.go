// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tests

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRedirectionHostDataSource(t *testing.T) {

	var id = rand.IntN(1e4)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + fmt.Sprintf(`
resource "nginxproxymanager_redirection_host" "test" {
  forward_domain_name = "test_%d.com"
  forward_http_code = 301
  forward_scheme = "https"
  domain_names = [
	"domain1_%d.com",
	"domain2_%d.com",
  ]

  preserve_path = false
}
data "nginxproxymanager_redirection_host" "test" {
  id = nginxproxymanager_redirection_host.test.id
}
`, id, id, id),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.nginxproxymanager_redirection_host.test", "forward_domain_name", fmt.Sprintf("test_%d.com", id)),
					resource.TestCheckResourceAttr("data.nginxproxymanager_redirection_host.test", "forward_http_code", "301"),
					resource.TestCheckResourceAttr("data.nginxproxymanager_redirection_host.test", "domain_names.#", "2"),
					resource.TestCheckResourceAttr("data.nginxproxymanager_redirection_host.test", "domain_names.0", fmt.Sprintf("domain1_%d.com", id)),
					resource.TestCheckResourceAttr("data.nginxproxymanager_redirection_host.test", "domain_names.1", fmt.Sprintf("domain2_%d.com", id)),
					resource.TestCheckResourceAttr("data.nginxproxymanager_redirection_host.test", "preserve_path", "false"),
				),
			},
		},
	})
}
