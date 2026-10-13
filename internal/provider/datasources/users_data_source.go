// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"strconv"

	models "github.com/edklesel/terraform-provider-nginxproxymanager/internal/models"
	sdk "github.com/edklesel/terraform-provider-nginxproxymanager/internal/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &usersDataSource{}
	_ datasource.DataSourceWithConfigure = &usersDataSource{}
)

func NewUsersDataSource() datasource.DataSource {
	return &usersDataSource{}
}

type usersDataSource struct {
	client *sdk.Client
}

// var userSchema

func (d *usersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_users"
}

func (d *usersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"users": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
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
				},
			},
		},
	}
}

type usersDataSourceModel struct {
	Users []models.UserModel `tfsdk:"users"`
}

func (d *usersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state usersDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	users, err := d.client.Users.List()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to list users.",
			err.Error(),
		)
		return
	}

	for _, user := range users {
		userState := models.UserModel{
			ID:         types.StringValue(strconv.Itoa(user.ID)),
			Email:      types.StringValue(*user.Email),
			Name:       types.StringValue(*user.Name),
			Nickname:   types.StringValue(*user.Nickname),
			IsDisabled: types.BoolValue(*user.IsDisabled),
			Avatar:     types.StringValue(user.Avatar),
		}

		for _, role := range *user.Roles {
			userState.Roles = append(userState.Roles, types.StringValue(role))
		}

		state.Users = append(state.Users, userState)
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *usersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	d.client = configureProvider(req, resp)
}
