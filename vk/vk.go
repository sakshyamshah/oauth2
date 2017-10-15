// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vk provides constants for using oidc to access VK.com.
package vk // import "github.com/sakshyamshah/oidc/vk"

import (
	"github.com/sakshyamshah/oidc"
)

// Endpoint is VK's OAuth 2.0 endpoint.
var Endpoint = oidc.Endpoint{
	AuthURL:  "https://oauth.vk.com/authorize",
	TokenURL: "https://oauth.vk.com/access_token",
}
