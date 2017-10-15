// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hipchat provides constants for using oidc to access HipChat.
package hipchat // import "github.com/sakshyamshah/oidc/hipchat"

import (
	"encoding/json"
	"errors"

	"github.com/sakshyamshah/oidc"
	"github.com/sakshyamshah/oidc/clientcredentials"
)

// Endpoint is HipChat's OAuth 2.0 endpoint.
var Endpoint = oidc.Endpoint{
	AuthURL:  "https://www.hipchat.com/users/authorize",
	TokenURL: "https://api.hipchat.com/v2/oauth/token",
}

// ServerEndpoint returns a new oidc.Endpoint for a HipChat Server instance
// running on the given domain or host.
func ServerEndpoint(host string) oidc.Endpoint {
	return oidc.Endpoint{
		AuthURL:  "https://" + host + "/users/authorize",
		TokenURL: "https://" + host + "/v2/oauth/token",
	}
}

// ClientCredentialsConfigFromCaps generates a Config from a HipChat API
// capabilities descriptor. It does not verify the scopes against the
// capabilities document at this time.
//
// For more information see: https://www.hipchat.com/docs/apiv2/method/get_capabilities
func ClientCredentialsConfigFromCaps(capsJSON []byte, clientID, clientSecret string, scopes ...string) (*clientcredentials.Config, error) {
	var caps struct {
		Caps struct {
			Endpoint struct {
				TokenURL string `json:"tokenUrl"`
			} `json:"oidcProvider"`
		} `json:"capabilities"`
	}

	if err := json.Unmarshal(capsJSON, &caps); err != nil {
		return nil, err
	}

	// Verify required fields.
	if caps.Caps.Endpoint.TokenURL == "" {
		return nil, errors.New("oidc/hipchat: missing oidc token URL in the capabilities descriptor JSON")
	}

	return &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		TokenURL:     caps.Caps.Endpoint.TokenURL,
	}, nil
}
