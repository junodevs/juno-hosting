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
	"github.com/junodevs/hosting-server/server/routes/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func init() {
	// Load Config
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	// Connect to Redis database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	auth.OAuthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/v1/callback",
		ClientID:     config.Config.OAuth.ClientID,
		ClientSecret: config.Config.OAuth.ClientSecret,
		Endpoint:     github.Endpoint,
		Scopes: []string{
			"user:email",
		},
	}

	// Start web server
	err := server.Start(config.Config.Port, config.Config.HostName)

	if err != nil {
		log.Fatal(err)
	}
}
