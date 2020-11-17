/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"log"

	"github.com/junodevs/hosting-server/config"
	"github.com/junodevs/hosting-server/database"
	"github.com/junodevs/hosting-server/server"
	"github.com/junodevs/hosting-server/server/v1/authentication"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func main() {
	// Load Config
	if err := config.Load(); err != nil {
		log.Fatalf("got error while loading config: %v", err)
	}

	// Connect to Redis database
	if err := database.Connect(); err != nil {
		log.Fatalf("got error when connecting to database: %v", err)
	}

	// Populate OAuth2 client configuration
	authentication.OAuthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/v1/callback",
		ClientID:     config.Config.OAuth.ClientID,
		ClientSecret: config.Config.OAuth.ClientSecret,
		Endpoint:     github.Endpoint,
		Scopes: []string{
			"user:email",
		},
	}

	// Start Fiber web server
	if _, err := server.Start(
		config.Config.Port,
		config.Config.HostName,
	); err != nil {
		log.Fatalf("could not start Fiber web server: %v", err)
	}
}
