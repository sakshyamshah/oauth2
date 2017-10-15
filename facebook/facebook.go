// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package facebook provides constants for using oidc to access Facebook.
package facebook // import "github.com/sakshyamshah/oidc/facebook"

import (
	"github.com/sakshyamshah/oidc"
)

// Endpoint is Facebook's OAuth 2.0 endpoint.
var Endpoint = oidc.Endpoint{
	AuthURL:  "https://www.facebook.com/dialog/oauth",
	TokenURL: "https://graph.facebook.com/oauth/access_token",
}
