/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/junodevs/hosting-server/server/v1/authentication"
)

// registerRoutes registers the API routes of the app
func registerRoutes(app *fiber.App) {
	v1 := app.Group("/v1")

	v1.Get("/me", authentication.MeRoute)
}
