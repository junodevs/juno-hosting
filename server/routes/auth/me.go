/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package auth

import "net/http"

// MeRoute represents the GET /me API route
func MeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
