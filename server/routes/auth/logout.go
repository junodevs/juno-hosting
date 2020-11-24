package auth

import (
	"github.com/markbates/goth/gothic"
	"net/http"
)

// LogoutRoute represents the GET /auth/logout API route
func LogoutRoute(w http.ResponseWriter, r *http.Request) {
	_ = gothic.Logout(w, r)
	http.Redirect(w, r, "/v1", http.StatusTemporaryRedirect)
}
