// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package uber provides constants for using oidc to access Uber.
package uber // import "github.com/sakshyamshah/oidc/uber"

import (
	"github.com/sakshyamshah/oidc"
)

// Endpoint is Uber's OAuth 2.0 endpoint.
var Endpoint = oidc.Endpoint{
	AuthURL:  "https://login.uber.com/oauth/v2/authorize",
	TokenURL: "https://login.uber.com/oauth/v2/token",
}
