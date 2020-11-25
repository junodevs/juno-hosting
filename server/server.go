/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/junodevs/juno-hosting/server/routes/auth"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/junodevs/juno-hosting/config"
)

// Start begins the web server on the port and hostname
func Start(port int, hostname string) error {
	r := chi.NewRouter()

	// Register HTTP middleware functions
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
		},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(middleware.AllowContentType(""))
	r.Use(httprate.LimitByIP(
		config.Config.RateLimit.Requests,
		config.Config.RateLimit.Duration*time.Second,
	))

	// Register API endpoints
	r.Route("/v1", func(r chi.Router) {
		r.Get("/callback", auth.CallbackRoute)
		r.Get("/login", auth.LoginRoute)
		r.Get("/logout", auth.LogoutRoute)
		r.Get("/me", auth.MeRoute)
	})

	fmt.Printf("Juno Hosting API server listening on %s:%d\n", hostname, port)

	return http.ListenAndServe(fmt.Sprintf("%s:%d", hostname, port), r)
}
