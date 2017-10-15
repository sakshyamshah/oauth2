// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package slack provides constants for using oidc to access Slack.
package slack // import "github.com/sakshyamshah/oidc/slack"

import (
	"github.com/sakshyamshah/oidc"
)

// Endpoint is Slack's OAuth 2.0 endpoint.
var Endpoint = oidc.Endpoint{
	AuthURL:  "https://slack.com/oauth/authorize",
	TokenURL: "https://slack.com/api/oauth.access",
}
