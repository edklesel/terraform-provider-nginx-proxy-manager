// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DeadHostModel struct {
	ID             types.String   `tfsdk:"id"`
	DomainNames    []types.String `tfsdk:"domain_names"`
	CertificateID  types.Int64    `tfsdk:"certificate_id"`
	HTTP2Support   types.Bool     `tfsdk:"http2_support"`
	HSTSEnabled    types.Bool     `tfsdk:"hsts_enabled"`
	HSTSSubdomains types.Bool     `tfsdk:"hsts_subdomains"`
	SSLForced      types.Bool     `tfsdk:"ssl_forced"`
	OwnerUserID    types.Int64    `tfsdk:"owner_user_id"`
	AdvancedConfig types.String   `tfsdk:"advanced_config"`
	Enabled        types.Bool     `tfsdk:"enabled"`
}
