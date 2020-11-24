/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package domain

// Response represents an API response
type Response struct {
	Status  int                    `json:"status"`
	Payload map[string]interface{} `json:"payload"`
	Error   string                 `json:"error"`
}
