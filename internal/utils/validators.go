// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RequireAttributesIfTrueValidator struct {
	RequiredAttrs []string
}

func (v RequireAttributesIfTrueValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Requires %v to be set if this attribute is true", strings.Join(v.RequiredAttrs, ", "))
}

func (v RequireAttributesIfTrueValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf("Requires `%v` to be set if this attribute is `true`", strings.Join(v.RequiredAttrs, "`, `"))
}

func (v RequireAttributesIfTrueValidator) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
	if req.ConfigValue.IsNull() || !req.ConfigValue.ValueBool() {
		return
	}

	for _, attr := range v.RequiredAttrs {
		var val types.String // Adjust type if needed
		diags := req.Config.GetAttribute(ctx, path.Root(attr), &val)
		resp.Diagnostics.Append(diags...)

		if diags.HasError() || val.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root(attr),
				"Missing required attribute",
				fmt.Sprintf("Attribute '%s' must be set when '%s' is true.", attr, req.Path.String()),
			)
		}
	}
}
