// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package yandex provides constants for using oidc to access Yandex APIs.
package yandex // import "github.com/sakshyamshah/oidc/yandex"

import (
	"github.com/sakshyamshah/oidc"
)

// Endpoint is the Yandex OAuth 2.0 endpoint.
var Endpoint = oidc.Endpoint{
	AuthURL:  "https://oauth.yandex.com/authorize",
	TokenURL: "https://oauth.yandex.com/token",
}
