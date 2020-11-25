/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package auth

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/markbates/goth/gothic"

	"github.com/junodevs/juno-hosting/domain"
)

// CallbackRoute represents the GET /auth/callback API route
func CallbackRoute(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		body, err := json.Marshal(&domain.Response{
			Status:  http.StatusInternalServerError,
			Payload: map[string]interface{}{},
			Error:   err.Error(),
		})

		if err != nil {
			log.Fatal(err)
		}

		_, _ = w.Write(body)

		return
	}

	t, _ := template.New("foo").Parse(userTemplate)
	_ = t.Execute(w, user)
}
