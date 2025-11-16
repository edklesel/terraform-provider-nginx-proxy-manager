// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type RedirectionHostModel struct {
	ID                types.String   `tfsdk:"id"`
	DomainNames       []types.String `tfsdk:"domain_names"`
	ForwardDomainName types.String   `tfsdk:"forward_domain_name"`
	ForwardHTTPCode   types.Int64    `tfsdk:"forward_http_code"`
	ForwardScheme     types.String   `tfsdk:"forward_scheme"`
	CertificateID     types.Int64    `tfsdk:"certificate_id"`
	SSLForced         types.Bool     `tfsdk:"ssl_forced"`
	BlockExploits     types.Bool     `tfsdk:"block_exploits"`
	AdvancedConfig    types.String   `tfsdk:"advanced_config"`
	HTTP2Support      types.Bool     `tfsdk:"http2_support"`
	Enabled           types.Bool     `tfsdk:"enabled"`
	HSTSEnabled       types.Bool     `tfsdk:"hsts_enabled"`
	HSTSSubdomains    types.Bool     `tfsdk:"hsts_subdomains"`
	PreservePath      types.Bool     `tfsdk:"preserve_path"`
	OwnerUserID       types.Int64    `tfsdk:"owner_user_id"`
}
