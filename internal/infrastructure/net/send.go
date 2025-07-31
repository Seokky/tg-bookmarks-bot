package net

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// PostBodyParams is map[string]any that describes POST request body payload
type PostBodyParams map[string]any

// Send do POST request and unmarshal response to R
func Send[R any](client Client, entrypoint string, params PostBodyParams) (R, error) {
	// Initialize R with zero value
	var formattedResult R

	// Build url string
	url, err := client.BuildURL(entrypoint)
	if err != nil {
		return formattedResult, err
	}

	// Marshall params
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return formattedResult, err
	}

	// Do request
	result, err := http.Post(url, "application/json", bytes.NewBuffer(paramsJSON))
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
