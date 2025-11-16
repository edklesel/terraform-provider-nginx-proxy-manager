// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package go_npm_proxy_manager

type DeadHostBase struct {
	DomainNames    *[]string `json:"domain_names,omitempty" client:"required"`
	AdvancedConfig *string   `json:"advanced_config,omitempty" proto:"# comment"`
	HTTP2Support   *bool     `json:"http2_support,omitempty" proto:"false"`
	HSTSEnabled    *bool     `json:"hsts_enabled,omitempty" proto:"false"`
	HSTSSubdomains *bool     `json:"hsts_subdomains,omitempty" proto:"false"`
	SSLForced      *bool     `json:"ssl_forced,omitempty" proto:"false"`
	CertificateID  *int      `json:"certificate_id,omitempty" proto:"0"`
}

type DeadHost struct {
	DeadHostBase
	ID          int       `json:"id"`
	CreatedOn   string    `json:"created_on" deep:"-"`
	ModifiedOn  string    `json:"modified_on" deep:"-"`
	OwnerUserID int       `json:"owner_user_id"`
	Meta        *HostMeta `json:"meta,omitempty" deep:"-"`
	Enabled     *bool     `json:"enabled"`
}
