package client

import (
	"net/url"

	"github.com/pciet/wichess"
)

// A sucessful Instance.Login puts a session key cookie into the jar.
func (an Instance) Login() error {
	l := url.Values{}
	l.Add(wichess.FormPlayerName, an.Name)
	l.Add(wichess.FormPassword, an.Password)
	return an.PostForm(wichess.LoginPath, l)
}
