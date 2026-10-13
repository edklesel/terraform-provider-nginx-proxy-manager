// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"

	sdk "github.com/edklesel/terraform-provider-nginxproxymanager/internal/sdk"

	"github.com/hashicorp/go-version"
)

var (
	// Certificate constraints.
	FieldLetsencryptEmail, _        = version.NewConstraint("< 2.13.0")
	ReturnDNSProviderCredentials, _ = version.NewConstraint("<2.13.0")
)

func CheckAttribute(r *sdk.Client, c *version.Constraints) bool {
	ver, err := r.GetVersion()
	if err != nil {
		panic(fmt.Errorf("unable to get client version: %s", err))
	}

	return c.Check(ver)
}
