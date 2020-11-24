/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package server

// Response represents an API response
type Response struct {
	Status  int      `json:"status"`
	Payload *Payload `json:"payload"`
	Error   string   `json:"error"`
}

// Payload is a shortcut to the map[string]interface{} type
type Payload map[string]interface{}
