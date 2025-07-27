package net

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type FetchQueryParam struct {
	Name, Value string
}

type FetchQueryParams []FetchQueryParam

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

// Attach query params to request
func attachQueryParams(req *http.Request, params FetchQueryParams) {
	q := req.URL.Query()

	for _, p := range params {
		q.Add(p.Name, p.Value)
	}

	req.URL.RawQuery = q.Encode()
}
