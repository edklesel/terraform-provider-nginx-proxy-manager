// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"

	sdk "github.com/edklesel/terraform-provider-nginxproxymanager/internal/sdk"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

func configureProvider(req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) (client *sdk.Client) {
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
