/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package authentication

import (
	"github.com/gofiber/fiber/v2"
	"github.com/junodevs/hosting-server/util"
	"golang.org/x/oauth2"
)

var (
	oauthStateString = util.RandomString(20)
	// OAuthConfig represents a oauth2 config
	OAuthConfig *oauth2.Config
)

// LoginRoute represents the POST /login API route
func LoginRoute(c *fiber.Ctx) error {
	return c.Redirect(
		OAuthConfig.AuthCodeURL(oauthStateString),
		fiber.StatusTemporaryRedirect,
	)
}
