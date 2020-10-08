package auth

import "net/http"

// SessionKeyCookie is the name of the cookie used to store the unique session key.
const SessionKeyCookie = "k"

// CreateBrowserSession sets the HttpOnly SessionKeyCookie. The key should have been encoded to
// base64 because only ASCII compatible characters are sent over HTTP.
func CreateBrowserSession(w http.ResponseWriter, key string) {
	http.SetCookie(w, &http.Cookie{
		Name:     SessionKeyCookie,
		Value:    key,
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
