/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package database

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/junodevs/hosting-server/config"
)

var (
	// Ctx is the context of the Redis connection
	Ctx = context.Background()

	// Redis is the active redis client connection
	Redis *redis.Client
)

// Connect creates a new connection to a Redis database
func Connect() error {
	uri, err := redis.ParseURL(config.Config.RedisURI)

	if err != nil {
		return err
	}

	Redis = redis.NewClient(uri)
	return nil
}
