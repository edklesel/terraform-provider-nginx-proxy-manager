// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tests

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccLetsencryptCertificateDataSource(t *testing.T) {

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
	resource "nginxproxymanager_letsencrypt_certificate" "test" {
	  domain_names = ["www2.example.com"]
	  dns_challenge = true
	  dns_provider = "powerdns"
	  dns_provider_credentials = "dns_powerdns_api_url = http://pdns:8081\r\ndns_powerdns_api_key = npm"
	  letsencrypt_email = "test@test.com"
	  propagation_seconds = 5
	}

	data "nginxproxymanager_letsencrypt_certificate" "test" {
	  id = nginxproxymanager_letsencrypt_certificate.test.id
	}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.nginxproxymanager_letsencrypt_certificate.test", "dns_provider", "powerdns"),
					resource.TestCheckResourceAttr("data.nginxproxymanager_letsencrypt_certificate.test", "dns_challenge", "true"),
					resource.TestCheckResourceAttr("data.nginxproxymanager_letsencrypt_certificate.test", "domain_names.#", "1"),
				),
			},
		},
	})
}
