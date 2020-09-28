package auth

import (
	"encoding/base64"
	"net/http"
)

// SessionKeyCookie is the name of the cookie used to store the unique session key.
const SessionKeyCookie = "k"

// CreateBrowserSession sets the HttpOnly SessionKeyCookie. The key string is encoded to base64
// so that only ASCII compatible characters are sent over HTTP.
func CreateBrowserSession(w http.ResponseWriter, sessionKey string) {
	http.SetCookie(w, &http.Cookie{
		Name:     SessionKeyCookie,
		Value:    base64.StdEncoding.EncodeToString([]byte(sessionKey)),
		Path:     "/",
		HttpOnly: true,
	})
}

// ClearBrowserSession invalidates the SessionKeyCookie and redirects to LoginPath.
func ClearBrowserSession(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     SessionKeyCookie,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
	http.Redirect(w, r, LoginPath, http.StatusSeeOther)
}
