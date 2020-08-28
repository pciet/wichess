package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (an Instance) Get(path string) error {
	_, err := an.JSONResponseGet(path)
	return err
}

// TODO: this is just a return body get, either check for JSON or change symbol

func (an Instance) JSONResponseGet(path string) (string, error) {
	resp, err := an.Client.Get(an.HTTPHost() + path)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad Get HTTP status %v on %v", resp.StatusCode, path)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
