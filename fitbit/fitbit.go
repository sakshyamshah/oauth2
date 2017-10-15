// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fitbit provides constants for using oidc to access the Fitbit API.
package fitbit // import "github.com/sakshyamshah/oidc/fitbit"

import (
	"github.com/sakshyamshah/oidc"
)

// Endpoint is the Fitbit API's OAuth 2.0 endpoint.
var Endpoint = oidc.Endpoint{
	AuthURL:  "https://www.fitbit.com/oidc/authorize",
	TokenURL: "https://api.fitbit.com/oidc/token",
}
