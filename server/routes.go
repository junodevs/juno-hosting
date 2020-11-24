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

func registerRoutes(app *fiber.App) {
	v1 := app.Group("/v1")

	v1.Get("/login", authentication.LoginRoute)
	v1.Get("/me", authentication.MeRoute)
	v1.Get("/callback", authentication.CallbackRoute)

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&Response{
			Status: 200,
			Payload: &Payload{
				"name":          "junohosting-server",
				"version":       "1.0.0",
				"project_url":   "https://github.com/junodevs/hosting-server",
				"documentation": "https://junodevs.github.io/hosting-server",
			},
			Error: "",
		})
	})
}
