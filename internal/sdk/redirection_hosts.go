// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package go_npm_proxy_manager

type RedirectionHostBase struct {
	ForwardScheme     *string   `json:"forward_scheme,omitempty" client:"required" proto:"http"`
	DomainNames       *[]string `json:"domain_names,omitempty" client:"required"`
	ForwardDomainName *string   `json:"forward_domain_name,omitempty" client:"required"`
	ForwardHTTPCode   *int      `json:"forward_http_code,omitempty" client:"required" proto:"301"`
	CertificateID     *int      `json:"certificate_id,omitempty"`
	SSLForced         *bool     `json:"ssl_forced,omitempty"`
	BlockExploits     *bool     `json:"block_exploits,omitempty"`
	AdvancedConfig    *string   `json:"advanced_config,omitempty"`
	HTTP2Support      *bool     `json:"http2_support,omitempty"`
	HSTSEnabled       *bool     `json:"hsts_enabled,omitempty"`
	HSTSSubdomains    *bool     `json:"hsts_subdomains,omitempty"`
	PreservePath      *bool     `json:"preserve_path,omitempty"`
}

type RedirectionHost struct {
	RedirectionHostBase
	ID          int       `json:"id"`
	CreatedOn   string    `json:"created_on" deep:"-"`
	ModifiedOn  string    `json:"modified_on" deep:"-"`
	OwnerUserID int       `json:"owner_user_id"`
	Meta        *HostMeta `json:"meta,omitempty" deep:"-"`
	Enabled     *bool     `json:"enabled"`
}
