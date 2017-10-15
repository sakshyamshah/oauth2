// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package linkedin provides constants for using oidc to access LinkedIn.
package linkedin // import "github.com/sakshyamshah/oidc/linkedin"

import (
	"github.com/sakshyamshah/oidc"
)

// Endpoint is LinkedIn's OAuth 2.0 endpoint.
var Endpoint = oidc.Endpoint{
	AuthURL:  "https://www.linkedin.com/uas/oidc/authorization",
	TokenURL: "https://www.linkedin.com/uas/oidc/accessToken",
}
