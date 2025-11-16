// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package go_npm_proxy_manager

type AccessListClient struct {
	Address   *string `json:"address,omitempty" proto:"10.1.2.3"`
	Directive *string `json:"directive,omitempty" proto:"allow"`
}

type AccessListItem struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty" deep:"-"` // Password isn't returned by the API
}

type AccessListBase struct {
	Name       *string             `json:"name,omitempty"`
	SatisfyAny *bool               `json:"satisfy_any,omitempty"`
	PassAuth   *bool               `json:"pass_auth,omitempty"`
	Items      *[]AccessListItem   `json:"items,omitempty" deep:"-"`   // Clients and Items aren't
	Clients    *[]AccessListClient `json:"clients,omitempty" deep:"-"` // returned by the GET method
}

type AccessListCreate struct {
	AccessListBase
}

type AccessList struct {
	ID          int    `json:"id"`
	CreatedOn   string `json:"created_on"`
	ModifiedOn  string `json:"modified_on"`
	OwnerUserID int    `json:"owner_user_id"`
	AccessListBase
	Meta      map[string]any `json:"meta,omitempty"`
	Directive *string        `json:"directive,omitempty" proto:"allow"`
}
