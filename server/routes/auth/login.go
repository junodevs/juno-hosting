/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/junodevs/juno-hosting/domain"
	"github.com/markbates/goth/gothic"
)

// LoginRoute represents the GET /auth/login API route
func LoginRoute(w http.ResponseWriter, r *http.Request) {
	// attempt to get user without re-auth
	if user, err := gothic.CompleteUserAuth(w, r); err == nil {
		body, err := json.Marshal(&domain.Response{
			Status: http.StatusOK,
			Payload: map[string]interface{}{
				"Name":        user.Name,
				"Email":       user.Email,
				"NickName":    user.NickName,
				"Location":    user.Location,
				"AvatarURL":   user.AvatarURL,
				"Description": user.Description,
				"UserID":      user.UserID,
			},
			Error: "",
		})

		if err != nil {
			log.Fatal(err)
		}

		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		w.Write(body)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}
