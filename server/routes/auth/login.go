/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package auth

import (
	"net/http"

	"github.com/junodevs/hosting-server/util"
	"golang.org/x/oauth2"
)

var (
	oauthStateString = util.RandomString(20)
	// OAuthConfig represents a oauth2 config
	OAuthConfig *oauth2.Config
)

// LoginRoute represents the GET /auth/login API route
func LoginRoute(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r,
		OAuthConfig.AuthCodeURL(oauthStateString),
		http.StatusTemporaryRedirect,
	)
}
