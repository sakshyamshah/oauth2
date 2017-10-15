// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package microsoft provides constants for using oidc to access Windows Live ID.
package microsoft // import "github.com/sakshyamshah/oidc/microsoft"

import (
	"github.com/sakshyamshah/oidc"
)

// LiveConnectEndpoint is Windows's Live ID OAuth 2.0 endpoint.
var LiveConnectEndpoint = oidc.Endpoint{
	AuthURL:  "https://login.live.com/oidc0_authorize.srf",
	TokenURL: "https://login.live.com/oidc0_token.srf",
}
