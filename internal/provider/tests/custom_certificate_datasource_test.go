// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tests

import (
	"encoding/base64"
	"fmt"
	"math/rand/v2"
	"testing"

	sdk "github.com/edklesel/terraform-provider-nginx-proxy-manager/internal/sdk"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCustomCertificateDataSource(t *testing.T) {

	id := rand.IntN(1e4)
	name := fmt.Sprintf("test_%d", id)
	cert, key := sdk.NewCertWithKey()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + fmt.Sprintf(`
	resource "nginxproxymanager_custom_certificate" "test" {
	  nice_name       = "%s"
	  certificate     = base64decode("%s")
	  certificate_key = base64decode("%s")
	}

	data "nginxproxymanager_custom_certificate" "test" {
		id = nginxproxymanager_custom_certificate.test.id
	}
`, name, base64.StdEncoding.EncodeToString([]byte(cert)), base64.StdEncoding.EncodeToString([]byte(key))),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.nginxproxymanager_custom_certificate.test", "nice_name", name),
					resource.TestCheckResourceAttr("data.nginxproxymanager_custom_certificate.test", "certificate", cert),
					resource.TestCheckResourceAttr("data.nginxproxymanager_custom_certificate.test", "certificate_key", key),
				),
			},
		},
	})
}
