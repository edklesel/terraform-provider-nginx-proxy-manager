// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package go_npm_proxy_manager

type HostMeta struct {
	NginxOnline *bool   `json:"nginx_online"`
	NginxErr    *string `json:"nginx_err"`
}
