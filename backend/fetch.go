package backend

import (
	"io"
	"net/http"
)

func (a *App) FetchURL(url string) (string, error) {
	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Convert the body to a string and return
	return string(body), nil
}
