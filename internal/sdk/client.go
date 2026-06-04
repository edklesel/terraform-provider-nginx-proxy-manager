// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package go_npm_proxy_manager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"

	"github.com/hashicorp/go-retryablehttp"
	version "github.com/hashicorp/go-version"
)

type Client struct {
	HostURL string     `json:"hostUrl"`
	Auth    AuthStruct `json:"auth"`
	Token   string
	Version *version.Version

	Users            *crudObject[User, UserBase, UserCreate]
	ProxyHosts       *crudObject[ProxyHost, ProxyHostBase, ProxyHostBase]
	RedirectionHosts *crudObject[RedirectionHost, RedirectionHostBase, RedirectionHostBase]
	DeadHosts        *crudObject[DeadHost, DeadHostBase, DeadHostBase]
	Streams          *crudObject[Stream, StreamBase, StreamBase]
	AccessLists      *crudObject[AccessList, AccessListBase, AccessListBase]
	Certificates     *certObject
}

type AuthStruct struct {
	Identity string `json:"identity"`
	Secret   string `json:"secret"`
}

type TokenResponseStruct struct {
	Token   string
	Expires string
}

type HTTPResponseError struct {
	Error HTTPResponseCode `json:"error"`
}

type HTTPResponseCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ClientVersion struct {
	Major    int64 `json:"major"`
	Minor    int64 `json:"minor"`
	Revision int64 `json:"revision"`
}

type ClientAPIRoot struct {
	Status  string        `json:"status"`
	Setup   bool          `json:"setup"`
	Version ClientVersion `json:"version"`
}

func CheckResponse(response []byte) *HTTPResponseError {
	var httperr HTTPResponseError

	// Check for boolean responses, which won't unmarshal properly
	if slices.Contains([]string{"true", "false"}, string(response)) {
		return nil
	}

	// If the response is not json, return a generic error code and the response
	if !json.Valid(response) {
		httperr = HTTPResponseError{
			Error: HTTPResponseCode{
				Code:    999,
				Message: string(response),
			},
		}
		return &httperr
	}

	// Otherwise if the response unmarshals successfully, then the
	// response is a json object (although it might not be an error)
	err := json.Unmarshal(response, &httperr)
	if err == nil {
		// If the Error.Code is 0, then the json returned wasn't an error and
		// we can return no error
		if httperr.Error.Code == 0 && httperr.Error.Message == "" {
			// if httperr.Error.Message == "" {
			return nil
		}
		// Otherwise, the response will be an error
		return &httperr

		// Lastly, if the json couldn't unmarshal properly, then the json body was
		// likely a list of items (i.e begins with "["), in which case there is
		// also no error to return
	} else {
		return nil
	}

}

func NewClient(url string, auth AuthStruct) *Client {

	var client = &Client{
		HostURL: url,
		Auth:    auth,
	}

	client.Users = &crudObject[User, UserBase, UserCreate]{
		client:          client,
		apiPath:         "users",
		writeOnlyFields: []string{},
	}
	client.ProxyHosts = &crudObject[ProxyHost, ProxyHostBase, ProxyHostBase]{
		client:          client,
		apiPath:         "nginx/proxy-hosts",
		writeOnlyFields: []string{},
	}
	client.RedirectionHosts = &crudObject[RedirectionHost, RedirectionHostBase, RedirectionHostBase]{
		client:          client,
		apiPath:         "nginx/redirection-hosts",
		writeOnlyFields: []string{},
	}
	client.DeadHosts = &crudObject[DeadHost, DeadHostBase, DeadHostBase]{
		client:          client,
		apiPath:         "nginx/dead-hosts",
		writeOnlyFields: []string{},
	}
	client.Streams = &crudObject[Stream, StreamBase, StreamBase]{
		client:          client,
		apiPath:         "nginx/streams",
		writeOnlyFields: []string{},
	}
	client.AccessLists = &crudObject[AccessList, AccessListBase, AccessListBase]{
		client:          client,
		apiPath:         "nginx/access-lists",
		writeOnlyFields: []string{},
	}
	client.Certificates = &certObject{
		client:          client,
		apiPath:         "nginx/certificates",
		writeOnlyFields: []string{"Meta.DNSProviderCredentials"},
	}

	_ver, err := client.GetVersion()
	if err != nil {
		panic(fmt.Sprintf("Unable to retrieve NPM version from API: %s", err))
	}
	client.Version = _ver

	return client
}

func (c *Client) GetToken() error {

	// Convert the username/password to json
	body, err := json.Marshal(c.Auth)
	if err != nil {
		return fmt.Errorf("error encoding client credentials to json: %s", err)
	}

	// Make the request for the token and parse the json response for the token value
	_response, err := c.NewRequest("tokens", http.MethodPost, &body, false, returnPointer("application/json"))
	if err != nil {
		return fmt.Errorf("unable to get client token: %s", err.Error())
	}

	var response TokenResponseStruct
	if err = json.Unmarshal(_response, &response); err != nil {
		panic(err.Error())
	}
	c.Token = response.Token

	return err
}

func (c Client) NewRequest(endpoint string, method string, body *[]byte, auth bool, content_type *string) (response []byte, err error) {

	// Format the URL based on the endpoint passed in
	url := fmt.Sprintf("%s/api/%s", c.HostURL, endpoint)

	// Structure the request
	var _body []byte
	if body != nil {
		_body = *body
	} else {
		_body = make([]byte, 0)
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(_body))
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}
	if body != nil {
		if content_type == nil {
			return nil, fmt.Errorf("Content-Type is required if a request body is provided")
		}
		req.Header.Add("Content-Type", *content_type)
	}

	// If auth is required, add the client token
	if auth {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	}

	// Make the request
	_client := retryablehttp.NewClient()
	_client.Logger = nil
	client := _client.StandardClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request to %s: %s", req.URL, err)
	}
	defer func() { _ = resp.Body.Close() }()

	// Read the response
	response, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error unwrapping response from %s: %s", req.URL, err)
	}

	// Check if the response is an error
	var httpErr = CheckResponse(response)
	if httpErr != nil {
		if httpErr.Error.Code == 999 || httpErr.Error.Message != "" {
			return nil, fmt.Errorf("error received from NPM API: %s", httpErr.Error.Message)
		} else {
			return nil, fmt.Errorf("error code %d received from NPM API: %s", httpErr.Error.Code, httpErr.Error.Message)
		}
	}

	// fmt.Println(string(response))

	return response, nil
}

func (c Client) GetVersion() (ver *version.Version, err error) {

	var response ClientAPIRoot
	_response, err := c.NewRequest("", http.MethodGet, nil, false, nil)
	if err != nil {
		return ver, fmt.Errorf("unable to get version of running NPM: %v", err)
	}

	err = json.Unmarshal(_response, &response)
	if err != nil {
		return
	}

	_ver, err := version.NewVersion(fmt.Sprintf("%d.%d.%d", response.Version.Major, response.Version.Minor, response.Version.Revision))
	if err != nil {
		return nil, err
	}

	return _ver, nil
}
