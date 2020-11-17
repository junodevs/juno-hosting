/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package config

import (
	"strings"
	"time"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
)

var (
	k = koanf.New(".")

	// Config represents a loaded config
	Config struct {
		HostName string `koanf:"hostname"`
		Port     int    `koanf:"port"`
		Prefork  bool   `koanf:"prefork"`
		RedisURI string `koanf:"redis_uri"`

		RateLimit struct {
			Requests int           `koanf:"requests"`
			Duration time.Duration `koanf:"duration"`
		} `koanf:"ratelimit"`

		OAuth struct {
			ClientID     string `koanf:"client_id"`
			ClientSecret string `koanf:"client_secret"`
		} `koanf:"oauth"`
	}
)

// Load configuration from environment or file
func Load() error {
	// 1. Load in configuration defaults
	k.Load(confmap.Provider(map[string]interface{}{
		"hostname":  "127.0.0.1",
		"port":      8080,
		"prefork":   false,
		"ratelimit": map[string]interface{}{},
		"redis_uri": "redis://127.0.0.1:6379/0",
		"oauth": map[string]interface{}{
			"client_id":     "",
			"client_secret": "",
		},
	}, "."), nil)

	// 2. Load configuration from JSON file
	if err := k.Load(file.Provider("./config.json"), json.Parser()); err != nil {
		return err
	}

	if err := k.Load(env.Provider("HOSTING_SERVER_", ".", func(s string) string {
		// Strip the prefix and replace any `_` with `.` so hierarchy is properly
		// represented.
		return strings.Replace(
			strings.ToLower(strings.TrimPrefix(s, "HOSTING_SERVER_")),
			"_",
			".",
			-1,
		)
	}), nil); err != nil {
		return err
	}

	if err := k.Unmarshal("", &Config); err != nil {
		return err
	}

	return nil
}
