/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/junodevs/hosting-server/config"
)

// Start begins the Fiber server on the port and hostname
func Start(port int, hostname string) (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		Prefork: config.Config.Prefork,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Default 500 status code
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				// Override status code if fiber.Error type
				code = e.Code
			}

			// Set Content-Type: text/plain; charset=utf-8
			c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

			// Return status code with error message
			return c.Status(code).JSON(&Response{
				Error:   err.Error(),
				Payload: &Payload{},
				Status:  code,
			})
		},
	})

	registerRoutes(app)
	registerMiddleware(app)

	return app, app.Listen(
		fmt.Sprintf("%s:%d", hostname, port),
	)
}
