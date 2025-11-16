// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type StreamModel struct {
	ID             types.String `tfsdk:"id"`
	IncomingPort   types.Int64  `tfsdk:"incoming_port"`
	TCPForwarding  types.Bool   `tfsdk:"tcp_forwarding"`
	UDPForwarding  types.Bool   `tfsdk:"udp_forwarding"`
	ForwardingHost types.String `tfsdk:"forwarding_host"`
	ForwardingPort types.Int64  `tfsdk:"forwarding_port"`
	OwnerUserID    types.Int64  `tfsdk:"owner_user_id"`
}
