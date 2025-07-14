package net

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

var host, base, token string

// Load environment variables
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("can't load .env file")
	}

	host = os.Getenv("TELEGRAM_HOST")
	base = os.Getenv("TELEGRAM_BASE")
	token = os.Getenv("TELEGRAM_API_TOKEN")
}

// Create client instance once
var client = http.Client{}

// Fetch do GET request and unmarshal response to R
func Fetch[R any](entrypoint string) (R, error) {
	// Initialize R with zero value
	var formattedResult R

	// Build url string
	url, err := buildUrl(entrypoint)
	if err != nil {
		return formattedResult, err
	}

	// TODO add cancel context
	// Create request with method and url
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return formattedResult, err
	}

	// Do request itself
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

func Send() {

}

// Build url string from host, base, token and entrypoint
func buildUrl(entrypoint string) (string, error) {
	if len(host) == 0 || len(base) == 0 || len(token) == 0 {
		return "", errors.New(fmt.Sprintf("can't build buildUrl with host = %s, base = %s, token = %s", host, base, token))
	}

	result := "https://" + path.Join(host, base+token, entrypoint)

	return result, nil
}
