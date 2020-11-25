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

		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		w.Write(body)

		return
	}

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
}
