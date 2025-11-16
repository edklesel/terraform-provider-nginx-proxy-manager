// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccUserDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
data "nginxproxymanager_user" "test" {
  id = 1
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.nginxproxymanager_user.test", "id", "1"),
					resource.TestCheckResourceAttr("data.nginxproxymanager_user.test", "name", "Administrator"),
					resource.TestCheckResourceAttr("data.nginxproxymanager_user.test", "nickname", "Admin"),
					resource.TestCheckResourceAttr("data.nginxproxymanager_user.test", "email", "admin@example.com"),
					resource.TestCheckResourceAttr("data.nginxproxymanager_user.test", "roles.0", "admin"),
				),
			},
		},
	})
}
