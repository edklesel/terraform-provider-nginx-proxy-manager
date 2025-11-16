// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type LetsencryptCertificateModel struct {
	ID                     types.String   `tfsdk:"id"`
	DomainNames            []types.String `tfsdk:"domain_names"`
	DNSChallenge           types.Bool     `tfsdk:"dns_challenge"`
	DNSProvider            types.String   `tfsdk:"dns_provider"`
	DNSProviderCredentials types.String   `tfsdk:"dns_provider_credentials"`
	LetsencryptEmail       types.String   `tfsdk:"letsencrypt_email"`
	PropagationSeconds     types.Int64    `tfsdk:"propagation_seconds"`
	NiceName               types.String   `tfsdk:"nice_name"`
	// From                   types.Int64    `tfsdk:"from"`
	// To                     types.Int64    `tfsdk:"to"`
	// CN                     types.String   `tfsdk:"cn"`
	// Issuer                 types.String   `tfsdk:"issuer"`
}

type CustomCertificateModel struct {
	ID                      types.String `tfsdk:"id"`
	NiceName                types.String `tfsdk:"nice_name"`
	Certificate             types.String `tfsdk:"certificate"`
	CertificateKey          types.String `tfsdk:"certificate_key"`
	IntermediateCertificate types.String `tfsdk:"intermediate_certificate"`
}
