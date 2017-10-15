// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package odnoklassniki provides constants for using oidc to access Odnoklassniki.
package odnoklassniki // import "github.com/sakshyamshah/oidc/odnoklassniki"

import (
	"github.com/sakshyamshah/oidc"
)

// Endpoint is Odnoklassniki's OAuth 2.0 endpoint.
var Endpoint = oidc.Endpoint{
	AuthURL:  "https://www.odnoklassniki.ru/oauth/authorize",
	TokenURL: "https://api.odnoklassniki.ru/oauth/token.do",
}
