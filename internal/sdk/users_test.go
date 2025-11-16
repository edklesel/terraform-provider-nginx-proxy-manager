// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package go_npm_proxy_manager

import (
	"testing"

	"github.com/edklesel/proto"
	"github.com/stretchr/testify/assert"
)

func TestUserCRUD(t *testing.T) {

	var client = newTestClient(t)

	testCRUD(t, client.Users)
}

func TestUser(t *testing.T) {

	client := newTestClient(t)

	t.Run("Log in as a new user", func(t *testing.T) {

		var _user UserCreate
		proto.Prototype(&_user)

		user, _ := client.Users.Create(_user)

		var auth UserCreateAuth
		proto.Prototype(&auth)

		err := client.Users.SetPassword(user.ID, auth)
		assert.NoError(t, err)

		var newAuth = AuthStruct{
			Identity: *_user.Email,
			Secret:   auth.Secret,
		}

		var newClient = NewClient(client.HostURL, newAuth)
		err = newClient.GetToken()
		if err != nil {
			t.Errorf("Failed creating test client: %v", err)
		}

		token := parseJWTToken(t, newClient.Token)

		assert.True(t, token.Valid)

		// Tidy up.
		err = client.Users.Delete(user.ID)
		if err != nil {
			t.Errorf("Unable to tidy up user: %v", err)
		}

	})

	t.Run("Update a user's permissions", func(t *testing.T) {

		var _user UserCreate
		proto.Prototype(&_user)

		user, _ := client.Users.Create(_user)

		var perms = Permissions{
			Visibility:       returnPointer("user"),
			AccessLists:      returnPointer("manage"),
			DeadHosts:        returnPointer("hidden"),
			ProxyHosts:       returnPointer("view"),
			RedirectionHosts: returnPointer("hidden"),
			Streams:          returnPointer("manage"),
			// Intentionally not updating certificate permissions
		}

		err := client.Users.UpdatePermissions(user.ID, perms)
		assert.NoError(t, err)

		// Re-fetch the user
		user2, err := client.Users.Get(user.ID)
		assert.NoError(t, err)

		// Check that the updated values match the expected
		assert.Equal(t, *perms.Visibility, *user2.Permissions.Visibility)
		assert.Equal(t, *perms.AccessLists, *user2.Permissions.AccessLists)
		assert.Equal(t, *perms.DeadHosts, *user2.Permissions.DeadHosts)
		assert.Equal(t, *perms.ProxyHosts, *user2.Permissions.ProxyHosts)
		assert.Equal(t, *perms.RedirectionHosts, *user2.Permissions.RedirectionHosts)
		assert.Equal(t, *perms.Streams, *user2.Permissions.Streams)

		// Check the Certificate value is the same as before
		assert.Equal(t, *user.Permissions.Certificates, *user2.Permissions.Certificates)

		// Tidy up
		err = client.Users.Delete(user.ID)
		if err != nil {
			t.Errorf("Unable to tidy up user: %v", err)
		}
	})

}
