// Package net works with net requests like GET and POST
package net

import (
	"errors"
	"net/http"
	"path"
)

var (
	ErrNoHostProvided     = errors.New("no host provided")
	ErrNoBaseProvided     = errors.New("no base provided")
	ErrNoTokenProvided    = errors.New("no token provided")
	ErrNoProtocolProvided = errors.New("no protocol provided")
)

// Client describes http client structure
type Client struct {
	client   *http.Client
	host     string
	base     string
	token    string
	protocol string
}

// NewClient returns pointer to [Client] instance
func NewClient(host, base, token, protocol string) *Client {
	return &Client{
		client:   &http.Client{},
		host:     host,
		base:     base,
		token:    token,
		protocol: protocol,
	}
}

// BuildURL performs constructing URL string from host, base, token and entrypoint
func (c Client) BuildURL(entrypoint string) (string, error) {
	if len(c.host) == 0 {
		return "", ErrNoHostProvided
	}

	if len(c.base) == 0 {
		return "", ErrNoBaseProvided
	}

	if len(c.token) == 0 {
		return "", ErrNoTokenProvided
	}

	if len(c.protocol) == 0 {
		return "", ErrNoProtocolProvided
	}

	result := c.protocol + "://" + path.Join(c.host, c.base+c.token, entrypoint)

	return result, nil
}
