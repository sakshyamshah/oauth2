// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package github provides constants for using oidc to access Github.
package github // import "github.com/sakshyamshah/oidc/github"

import (
	"github.com/sakshyamshah/oidc"
)

// Endpoint is Github's OAuth 2.0 endpoint.
var Endpoint = oidc.Endpoint{
	AuthURL:  "https://github.com/login/oauth/authorize",
	TokenURL: "https://github.com/login/oauth/access_token",
}
