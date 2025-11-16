// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"strconv"

	models "github.com/edklesel/terraform-provider-nginx-proxy-manager/internal/models"
	sdk "github.com/edklesel/terraform-provider-nginx-proxy-manager/internal/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &userDataSource{}
	_ datasource.DataSourceWithConfigure = &userDataSource{}
)

func NewUserDataSource() datasource.DataSource {
	return &userDataSource{}
}

type userDataSource struct {
	client *sdk.Client
}

// var userSchema

func (d *userDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (d *userDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:    true,
				Description: "ID of the user (required).",
			},
			"email": schema.StringAttribute{
				Computed:    true,
				Description: "Email of the user, used for logging in.",
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: "Username of the user.",
			},
			"nickname": schema.StringAttribute{
				Computed:    true,
				Description: "Nickname for the user.",
			},
			"is_disabled": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether the user is disabled or not.",
			},
			"roles": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "A list of roles that the user has.",
			},
			"avatar": schema.StringAttribute{
				Computed:    true,
				Description: "URL to the avatar for the user.",
			},
		},
	}
}

func (d *userDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var state models.UserModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	id, err := strconv.Atoi(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to convert ID value %s to int type", state.ID),
			err.Error(),
		)
		return
	}

	user, err := d.client.Users.Get(id)
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read user with ID %s.", state.ID.ValueString()),
			err.Error(),
		)
		return
	}

	state.ID = types.StringValue(strconv.Itoa(user.ID))
	state.Email = types.StringValue(*user.Email)
	state.Name = types.StringValue(*user.Name)
	state.Nickname = types.StringValue(*user.Nickname)
	state.IsDisabled = types.BoolValue(*user.IsDisabled)
	state.Avatar = types.StringValue(user.Avatar)

	state.Roles = []types.String{}
	for _, role := range *user.Roles {
		state.Roles = append(state.Roles, types.StringValue(role))
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *userDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = configureProvider(req, resp)
}
