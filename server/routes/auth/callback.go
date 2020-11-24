/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/junodevs/hosting-server/domain"
	"golang.org/x/oauth2"
)

func getUserInfo(state, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := OAuthConfig.Exchange(oauth2.NoContext, code)

	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %v", err)
	}

	request, err := http.NewRequest("GET", "https://api.github.com/user", nil)

	request.Header.Set("Authorization", "token "+token.AccessToken)

	if err != nil {
		return nil, fmt.Errorf("failed retrieving user info: %v", err)
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, fmt.Errorf("failed retrieving user info: %v", err)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %v", err)
	}

	return contents, nil
}

// CallbackRoute represents the GET /auth/callback API route
func CallbackRoute(w http.ResponseWriter, r *http.Request) {
	userInfo, err := getUserInfo(
		chi.URLParam(r, "state"),
		chi.URLParam(r, "code"),
	)

	if err != nil {
		body, err := json.Marshal(&domain.Response{
			Status:  http.StatusInternalServerError,
			Payload: map[string]interface{}{},
			Error:   err.Error(),
		})

		if err != nil {
			log.Fatal(err)
		}

		w.Write(body)
	}

	fmt.Println(string(userInfo))
	http.Redirect(w, r, "/v1", http.StatusMovedPermanently)
}
