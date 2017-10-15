// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mediamath provides constants for using oidc to access MediaMath.
package mediamath // import "github.com/sakshyamshah/oidc/mediamath"

import (
	"github.com/sakshyamshah/oidc"
)

// Endpoint is MediaMath's OAuth 2.0 endpoint for production.
var Endpoint = oidc.Endpoint{
	AuthURL:  "https://api.mediamath.com/oidc/v1.0/authorize",
	TokenURL: "https://api.mediamath.com/oidc/v1.0/token",
}

// SandboxEndpoint is MediaMath's OAuth 2.0 endpoint for sandbox.
var SandboxEndpoint = oidc.Endpoint{
	AuthURL:  "https://t1sandbox.mediamath.com/oidc/v1.0/authorize",
	TokenURL: "https://t1sandbox.mediamath.com/oidc/v1.0/token",
}
