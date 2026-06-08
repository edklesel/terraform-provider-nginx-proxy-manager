// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tests

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDeadHostDataSource(t *testing.T) {

	var id = rand.IntN(1e4)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + fmt.Sprintf(`
resource "nginxproxymanager_dead_host" "test" {
  domain_names = [
	"domain1_%d.com",
	"domain2_%d.com",
  ]
}

data "nginxproxymanager_dead_host" "test" {
  id = nginxproxymanager_dead_host.test.id
}
`, id, id),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.nginxproxymanager_dead_host.test", "domain_names.#", "2"),
					resource.TestCheckResourceAttr("data.nginxproxymanager_dead_host.test", "domain_names.0", fmt.Sprintf("domain1_%d.com", id)),
					resource.TestCheckResourceAttr("data.nginxproxymanager_dead_host.test", "domain_names.1", fmt.Sprintf("domain2_%d.com", id)),
				),
			},
		},
	})
}
