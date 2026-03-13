// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package go_npm_proxy_manager

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"testing"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func newTestClient(t *testing.T) (client *Client) {

	var auth = AuthStruct{
		"admin@example.com",
		"changeme",
	}

	client = NewClient("http://localhost:81", auth)

	err := client.GetToken()
	assert.NoError(t, err)
	return client

}

type keys struct {
	Key string `json:"key"`
	Pub string `json:"pub"`
}

func getKeys() (keys, error) {

	var keys keys

	_keys, err := os.ReadFile("../../docker/npm/keys.json")
	if err != nil {
		return keys, fmt.Errorf("unable to read keys file.")
	}

	err = json.Unmarshal(_keys, &keys)
	if err != nil {
		return keys, fmt.Errorf("unable to unmarshal keys file.")
	}

	_, err = jwt.ParseRSAPublicKeyFromPEM([]byte(keys.Pub))
	if err != nil {
		return keys, fmt.Errorf("unable to read public key format.")
	}

	_, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(keys.Key))
	if err != nil {
		return keys, fmt.Errorf("unable to read private key format.")
	}

	return keys, nil
}

func parseJWTToken(t *testing.T, clientToken string) *jwt.Token {

	// Verify that it's a valid JWT using the key files,
	// meaning the token was returned correctly
	jwtToken, _ := jwt.Parse(clientToken, func(token *jwt.Token) (interface{}, error) {
		var keys keys
		keys, err := getKeys()
		if err != nil {
			t.Errorf("unable to read npm keys.json")
			t.FailNow()
		}
		key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(keys.Pub))
		if err != nil {
			t.Errorf("unable to parse public key")
			t.FailNow()
		}
		return key, nil
	})

	return jwtToken
}

func TestGetToken(t *testing.T) {

	var client = newTestClient(t)

	t.Run("test fetching a token", func(t *testing.T) {

		// Fetch the token
		err := client.GetToken()
		if err != nil {
			t.Fatalf("%s", err)
		}

		// Verify that it's a valid JWT using the key files,
		// meaning the token was returned correctly
		token := parseJWTToken(t, client.Token)

		assert.True(t, token.Valid)
	})
}

func TestNewRequest(t *testing.T) {

	// Get the current version
	var client = newTestClient(t)

	t.Run("test API endpoint", func(t *testing.T) {
		resp, err := client.NewRequest("", http.MethodGet, nil, false, nil)
		if err != nil {
			t.Errorf("Error received from api: %s", err)
			t.FailNow()
		}

		var api ClientAPIRoot
		err = json.Unmarshal(resp, &api)
		if err != nil {
			t.Errorf("unable to unmarshal response from /api: %s", err)
		}

		assert.Equal(t, "OK", api.Status)
	})

}

func TestVersion(t *testing.T) {

	var client = newTestClient(t)

	t.Run("test version retrievable", func(t *testing.T) {

		ver, err := client.GetVersion()

		assert.NoError(t, err)
		assert.Regexp(t, regexp.MustCompile(`\d+\.\d+\.\d+`), ver.String())
	})
}