// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package go_npm_proxy_manager

type StreamBase struct {
	IncomingPort   *int    `json:"incoming_port,omitempty" proto:"8443"`
	ForwardingHost *string `json:"forwarding_host,omitempty"`
	ForwardingPort *int    `json:"forwarding_port,omitempty" proto:"8443"`
	TCPForwarding  *bool   `json:"tcp_forwarding,omitempty"`
	UDPForwarding  *bool   `json:"udp_forwarding,omitempty"`
}

type Stream struct {
	StreamBase
	ID          int       `json:"id"`
	CreatedOn   string    `json:"created_on" deep:"-"`
	ModifiedOn  string    `json:"modified_on" deep:"-"`
	OwnerUserID int       `json:"owner_user_id"`
	Meta        *HostMeta `json:"meta,omitempty" deep:"-"`
	Enabled     *bool     `json:"enabled"`
}
