// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foursquare provides constants for using oidc to access Foursquare.
package foursquare // import "github.com/sakshyamshah/oidc/foursquare"

import (
	"github.com/sakshyamshah/oidc"
)

// Endpoint is Foursquare's OAuth 2.0 endpoint.
var Endpoint = oidc.Endpoint{
	AuthURL:  "https://foursquare.com/oidc/authorize",
	TokenURL: "https://foursquare.com/oidc/access_token",
}
