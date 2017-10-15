// Copyright 2015 The oidc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bitbucket provides constants for using oidc to access Bitbucket.
package bitbucket

import (
	"github.com/sakshyamshah/oidc"
)

// Endpoint is Bitbucket's OAuth 2.0 endpoint.
var Endpoint = oidc.Endpoint{
	AuthURL:  "https://bitbucket.org/site/oidc/authorize",
	TokenURL: "https://bitbucket.org/site/oidc/access_token",
}
