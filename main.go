/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"fmt"
	"log"

	"github.com/markbates/goth/providers/discord"

	"github.com/junodevs/juno-hosting/config"
	"github.com/junodevs/juno-hosting/database"
	"github.com/junodevs/juno-hosting/server"

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
		github.New(config.Config.GithubOAuth.ClientID, config.Config.GithubOAuth.ClientSecret,
			fmt.Sprintf("http://%s:%d/v1/callback?provider=github",
				config.Config.HostName, config.Config.Port),
			"email"),

		discord.New(config.Config.DiscordOAuth.ClientID, config.Config.DiscordOAuth.ClientSecret,
			fmt.Sprintf("http://%s:%d/v1/callback?provider=discord",
				config.Config.HostName, config.Config.Port),
			discord.ScopeIdentify, discord.ScopeEmail),
	)

	// Start web server
	err := server.Start(config.Config.Port, config.Config.HostName)

	if err != nil {
		log.Fatal(err)
	}
}
