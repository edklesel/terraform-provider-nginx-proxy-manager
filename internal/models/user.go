// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type UserModel struct {
	ID         types.String   `tfsdk:"id"`
	Email      types.String   `tfsdk:"email"`
	Name       types.String   `tfsdk:"name"`
	Nickname   types.String   `tfsdk:"nickname"`
	IsDisabled types.Bool     `tfsdk:"is_disabled"`
	Roles      []types.String `tfsdk:"roles"`
	Avatar     types.String   `tfsdk:"avatar"`
}

type UserPasswordModel struct {
	UserID   types.Int64  `tfsdk:"user_id"`
	Password types.String `tfsdk:"password"`
}

type UserPermissionModel struct {
	Visibility       types.String `tfsdk:"visibility"`
	ProxyHosts       types.String `tfsdk:"proxy_hosts"`
	RedirectionHosts types.String `tfsdk:"redirection_hosts"`
	DeadHosts        types.String `tfsdk:"dead_hosts"`
	Streams          types.String `tfsdk:"streams"`
	AccessLists      types.String `tfsdk:"access_lists"`
	Certificates     types.String `tfsdk:"certificates"`
}

type UserPermissionsModel struct {
	UserID      types.Int64         `tfsdk:"user_id"`
	Permissions UserPermissionModel `tfsdk:"permissions"`
}
