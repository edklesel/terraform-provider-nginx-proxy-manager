// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"

	datasources "github.com/edklesel/terraform-provider-nginx-proxy-manager/internal/provider/datasources"
	resources "github.com/edklesel/terraform-provider-nginx-proxy-manager/internal/provider/resources"
	sdk "github.com/edklesel/terraform-provider-nginx-proxy-manager/internal/sdk"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure npmProvider satisfies various provider interfaces.
var (
	_ provider.Provider                       = &npmProvider{}
	_ provider.ProviderWithFunctions          = &npmProvider{}
	_ provider.ProviderWithEphemeralResources = &npmProvider{}

	minVersion = "2.15.0"
)

// npmProvider defines the provider implementation.
type npmProvider struct {
	version string
}

// npmProviderModel describes the provider data model.
type npmProviderModel struct {
	Host     types.String `tfsdk:"host"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
	Secure   types.Bool   `tfsdk:"secure"`
}

func (p *npmProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "nginxproxymanager"
	resp.Version = p.version
}

func (p *npmProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				MarkdownDescription: "Hostname of Nginx Proxy Manager host",
				Required:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "Email address of the user to authenticate as",
				Required:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password of the user to authenticate as",
				Required:            true,
				Sensitive:           true,
			},
			"secure": schema.BoolAttribute{
				MarkdownDescription: "Whether to use a secure (https) connection for API requests",
				Optional:            true,
			},
		},
	}
}

func (p *npmProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	tflog.Info(ctx, "Configuring Nginx Proxy Manager client.")

	var data npmProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown Nginx Proxy Manager Host",
			"The 'host' parameter is required for the Nginx Proxy Manager provider.",
		)
	}

	if data.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown Nginx Proxy Manager Username",
			"The 'username' parameter is required for the Nginx Proxy Manager provider.",
		)
	}

	if data.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown Nginx Proxy Manager Password",
			"The 'password' parameter is required for the Nginx Proxy Manager provider.",
		)
	}

	var protocol = "http"
	if data.Secure.ValueBool() {
		protocol = "https"
	}

	ctx = tflog.SetField(ctx, "nginx_proxy_manager_host", data.Host.ValueString())
	ctx = tflog.SetField(ctx, "nginx_proxy_manager_username", data.Username.ValueString())
	ctx = tflog.SetField(ctx, "nginx_proxy_manager_password", data.Password.ValueString())
	ctx = tflog.SetField(ctx, "nginx_proxy_manager_secure", data.Secure.ValueBool())

	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "nginx_proxy_manager_password")

	tflog.Debug(ctx, "Creating Nginx Proxy Manager client.")

	client := sdk.NewClient(
		fmt.Sprintf("%s://%s", protocol, data.Host.ValueString()),
		sdk.AuthStruct{
			Identity: data.Username.ValueString(),
			Secret:   data.Password.ValueString(),
		},
	)

	// Check if NPM is the minimum version
	versionConstraint, err := version.NewConstraint(fmt.Sprintf(">= %s", minVersion))
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed checking the minimum version",
			fmt.Sprintf(
				"The following error occurred while checking the minimum version: %s",
				err.Error(),
			),
		)
	}
	if !versionConstraint.Check(client.Version) {
		resp.Diagnostics.AddError(
			"Minimum version requirements not met",
			fmt.Sprintf(
				"The NGINX Proxy Manager Terraform Provider requires version %s or above.",
				minVersion,
			),
		)
		return
	}

	tflog.Info(ctx, "Client created, fetching token.")

	err = client.GetToken()
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed creating Nginx Proxy Manager client",
			fmt.Sprintf(
				"The following error occurred while creating the Nginx Proxy Manager client: %s",
				err.Error(),
			),
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Configured Nginx Proxy Manager client.", map[string]any{"success": true})

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *npmProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		resources.NewUserResource,
		resources.NewUserPasswordResource,
		resources.NewUserPermissionsResource,
		resources.NewAccessListResource,
		resources.NewProxyHostResource,
		resources.NewDeadHostResource,
		resources.NewRedirectionHostResource,
		resources.NewStreamResource,
		resources.NewCustomCertificateResource,
		resources.NewLetsencryptCertificateResource,
	}
}

func (p *npmProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{}
}

func (p *npmProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		datasources.NewUserDataSource,
		datasources.NewUsersDataSource,
		datasources.NewProxyHostDataSource,
		datasources.NewAccessListDataSource,
		datasources.NewDeadHostDataSource,
		datasources.NewRedirectionHostDataSource,
		datasources.NewStreamDataSource,
		datasources.NewCustomCertificateDataSource,
		datasources.NewletsencryptCertificateDataSource,
	}
}

func (p *npmProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &npmProvider{
			version: version,
		}
	}
}
