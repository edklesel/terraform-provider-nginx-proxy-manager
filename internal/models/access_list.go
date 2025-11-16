// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessListItemModel struct {
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

type AccessListClientModel struct {
	Address   types.String `tfsdk:"address"`
	Directive types.String `tfsdk:"directive"`
}

type AccessListModel struct {
	ID          types.String            `tfsdk:"id"`
	Name        types.String            `tfsdk:"name"`
	SatisfyAny  types.Bool              `tfsdk:"satisfy_any"`
	PassAuth    types.Bool              `tfsdk:"pass_auth"`
	Items       []AccessListItemModel   `tfsdk:"item"`
	Clients     []AccessListClientModel `tfsdk:"client"`
	OwnerUserID types.Int64             `tfsdk:"owner_user_id"`
}
