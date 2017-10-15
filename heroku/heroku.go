// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package heroku provides constants for using oidc to access Heroku.
package heroku // import "github.com/sakshyamshah/oidc/heroku"

import (
	"github.com/sakshyamshah/oidc"
)

// Endpoint is Heroku's OAuth 2.0 endpoint.
var Endpoint = oidc.Endpoint{
	AuthURL:  "https://id.heroku.com/oauth/authorize",
	TokenURL: "https://id.heroku.com/oauth/token",
}
