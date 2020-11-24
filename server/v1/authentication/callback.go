/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package authentication

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
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

// CallbackRoute represents the POST /login API route
func CallbackRoute(c *fiber.Ctx) error {
	userInfo, err := getUserInfo(c.Query("state"), c.Query("code"))

	if err != nil {
		return err
	}

	fmt.Println(string(userInfo))
	c.Redirect("/v1/")

	return nil
}
