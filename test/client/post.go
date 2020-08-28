package client

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// JSONResponsePost expects a JSON response to the HTTP Post request. The JSON is returned
// undecoded as a string.
func (an Instance) JSONResponsePost(path, contentType string,
	requestBody io.Reader) (string, error) {

	resp, err := an.Client.Post(an.HTTPHost()+path, contentType, requestBody)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad Post HTTP status %v on %v", resp.StatusCode, path)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (an Instance) Post(path string, contentType string, body io.Reader) error {
	resp, err := an.Client.Post(an.HTTPHost()+path, contentType, body)
	if err != nil {
		return err
	}
	_, err = io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		return err
	}
	err = resp.Body.Close()
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad Post HTTP status %v on %v", resp.StatusCode, path)
	}
	return nil
}

func (an Instance) PostForm(path string, fields url.Values) error {
	resp, err := an.Client.PostForm(an.HTTPHost()+path, fields)
	if err != nil {
		return err
	}
	_, err = io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		return err
	}
	err = resp.Body.Close()
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad PostForm HTTP status %v on %v", resp.StatusCode, path)
	}
	return nil
}
