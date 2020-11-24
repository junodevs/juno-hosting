/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"fmt"
	"log"

	"github.com/junodevs/hosting-server/config"
	"github.com/junodevs/hosting-server/database"
	"github.com/junodevs/hosting-server/server"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
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

	goth.UseProviders(
		github.New(config.Config.OAuth.ClientID, config.Config.OAuth.ClientSecret, fmt.Sprintf("http://%s:%d/v1/callback?provider=github", config.Config.HostName, config.Config.Port), "user", "email"),
	)

	// Start web server
	err := server.Start(config.Config.Port, config.Config.HostName)

	if err != nil {
		log.Fatal(err)
	}
}
