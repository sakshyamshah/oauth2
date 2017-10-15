// Copyright 2017 The oidc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package amazon provides constants for using oidc to access Amazon.
package amazon

import (
	"github.com/sakshyamshah/oidc"
)

// Endpoint is Amazon's OAuth 2.0 endpoint.
var Endpoint = oidc.Endpoint{
	AuthURL:  "https://www.amazon.com/ap/oa",
	TokenURL: "https://api.amazon.com/auth/o2/token",
}
