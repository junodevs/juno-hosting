/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/junodevs/hosting-server/config"
)

// registerMiddleware registers the Fiber middleware used in the app
func registerMiddleware(app *fiber.App) {
	app.Use(limiter.New(limiter.Config{
		Duration: config.Config.RateLimit.Duration,
		Max:      config.Config.RateLimit.Requests,
	}))

	app.Use(cors.New())
	app.Use(logger.New())

	// Custom middleware to set security-related headers
	app.Use(func(c *fiber.Ctx) error {
		// Set some security headers
		c.Set("X-Download-Options", "noopen")
		c.Set("X-DNS-Prefetch-Control", "off")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("Referrer-Policy", "no-referrer-when-downgrade")
		c.Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'; base-uri 'none'; form-action 'none';")

		// Go to next middleware
		return c.Next()
	})
}
