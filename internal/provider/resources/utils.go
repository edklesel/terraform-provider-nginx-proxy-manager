// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	sdk "github.com/edklesel/terraform-provider-nginx-proxy-manager/internal/sdk"
)

func configureProvider(req resource.ConfigureRequest, resp *resource.ConfigureResponse) (client *sdk.Client) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected configuration type",
			fmt.Sprintf("Expected *sdk.Client, got: %T. Please report this issue to the provider developer.", req.ProviderData),
		)
		return
	}

	return client
}
