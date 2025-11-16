// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package go_npm_proxy_manager

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type UserCreateAuth struct {
	Type   string `json:"type" proto:"password"`
	Secret string `json:"secret" proto:"123456789abcdef"`
}

type UserBase struct {
	Email      *string   `json:"email,omitempty" client:"required" proto:"example@test.com"`
	Name       *string   `json:"name,omitempty" client:"required"`
	Nickname   *string   `json:"nickname,omitempty" client:"required"`
	IsDisabled *bool     `json:"is_disabled,omitempty" proto:"false"`
	Roles      *[]string `json:"roles,omitempty" deep:"-" proto:"admin"`
}

type UserCreate struct {
	UserBase
}

type User struct {
	UserBase
	ID          int         `json:"id"`
	Permissions Permissions `json:"permissions"`
	CreatedOn   string      `json:"created_on"`
	ModifiedOn  string      `json:"modified_on"`
	Avatar      string      `json:"avatar"`
}

type Permissions struct {
	Visibility       *string `json:"visibility,omitempty"`
	ProxyHosts       *string `json:"proxy_hosts,omitempty"`
	RedirectionHosts *string `json:"redirection_hosts,omitempty"`
	DeadHosts        *string `json:"dead_hosts,omitempty"`
	Streams          *string `json:"streams,omitempty"`
	AccessLists      *string `json:"access_lists,omitempty"`
	Certificates     *string `json:"certificates,omitempty"`
}

func (o crudObject[User, UserBase, UserCreate]) SetPassword(id int, auth UserCreateAuth) (err error) {

	body, err := json.Marshal(auth)
	if err != nil {
		return fmt.Errorf("unable to marshal provided UserCreateAuth object into json: %s", err)
	}

	_response, err := o.client.NewRequest(
		fmt.Sprintf("%s/%d/auth", o.apiPath, id),
		http.MethodPut,
		&body,
		true,
		returnPointer("application/json"),
	)
	if err != nil {
		return fmt.Errorf("error setting password for user %d: %s", id, err)
	}

	response, err := strconv.ParseBool(string(_response))
	if err != nil {
		return fmt.Errorf("error parsing response while setting password for user %d. Expected bool, got: %s", id, err)
	} else if !response {
		return fmt.Errorf("failed to set password for user %d", id)
	}

	return err

}

func (o crudObject[User, UserBase, UserCreate]) UpdatePermissions(id int, permissions Permissions) (err error) {

	body, err := json.Marshal(permissions)
	if err != nil {
		return fmt.Errorf("unable to marshal provided Permissions object into json: %s", err)
	}

	_response, err := o.client.NewRequest(
		fmt.Sprintf("%s/%d/permissions", o.apiPath, id),
		http.MethodPut,
		&body,
		true,
		returnPointer("application/json"),
	)
	if err != nil {
		return fmt.Errorf("error updating permissions for user %d: %s", id, err)
	}

	response, err := strconv.ParseBool(string(_response))
	if err != nil {
		return fmt.Errorf("error parsing response while updating permissions for user %d. Expected bool, got: %s", id, err)
	} else if !response {
		return fmt.Errorf("failed to update permissions for user %d", id)
	}

	return err
}
