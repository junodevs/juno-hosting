/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package auth

import (
	"html/template"
	"net/http"

	"github.com/markbates/goth/gothic"
)

// LoginRoute represents the GET /auth/login API route
func LoginRoute(w http.ResponseWriter, r *http.Request) {
	// attempt to get user without re-auth
	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		t, _ := template.New("foo").Parse(userTemplate)
		_ = t.Execute(w, gothUser)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}
