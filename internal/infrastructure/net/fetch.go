package net

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
)

// Fetch do GET request and unmarshal response to R
func Fetch[R any](entrypoint string, params FetchQueryParams) (R, error) {
	// Initialize R with zero value
	var formattedResult R

	// Build url string
	url, err := buildURL(entrypoint)
	if err != nil {
		return formattedResult, err
	}

	// Create request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return formattedResult, err
	}

	// Attach query params to request
	attachQueryParams(req, params)

	fmt.Println("poll()", req.URL.String())

	// Do request
	result, err := client.Do(req)
	if err != nil {
		return formattedResult, err
	}
	defer func() { _ = result.Body.Close() }()

	// Handle unwanted status codes
	if result.StatusCode != http.StatusOK {
		return formattedResult, errors.New(result.Status)
	}

	// Read response body as bytes
	data, err := io.ReadAll(result.Body)
	if err != nil {
		return formattedResult, err
	}

	// Unmarshall data to R
	err = json.Unmarshal(data, &formattedResult)
	if err != nil {
		return formattedResult, err
	}

	return formattedResult, nil
}

// Build url string from host, base, token and entrypoint
func buildURL(entrypoint string) (string, error) {
	if len(host) == 0 || len(base) == 0 || len(token) == 0 {
		return "", fmt.Errorf("can't build buildURL with host = %s, base = %s, token = %s", host, base, token)
	}

	result := "https://" + path.Join(host, base+token, entrypoint)

	return result, nil
}

// Attach query params to request
func attachQueryParams(req *http.Request, params FetchQueryParams) {
	q := req.URL.Query()

	for _, p := range params {
		q.Add(p.Name, p.Value)
	}

	req.URL.RawQuery = q.Encode()
}

type FetchQueryParam struct {
	Name, Value string
}

type FetchQueryParams []FetchQueryParam
