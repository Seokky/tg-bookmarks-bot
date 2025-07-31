package net_test

import (
	"errors"
	"fmt"
	"testing"
	"tg-bookmarks-bot/internal/infrastructure/env"
	"tg-bookmarks-bot/internal/infrastructure/net"
)

func Test_buildURL(t *testing.T) {
	t.Parallel()

	cases := []struct {
		client      *net.Client
		entrypoint  string
		expected    string
		expectedErr error
	}{
		{
			client: net.NewClient(
				env.Get("TEST_HOST"),
				env.Get("TEST_BASE"),
				env.Get("TEST_TOKEN"),
				"https",
			),
			entrypoint: "entrypoint",
			expected: fmt.Sprintf(
				"https://%s/%s%s/entrypoint",
				env.Get("TEST_HOST"),
				env.Get("TEST_BASE"),
				env.Get("TEST_TOKEN"),
			),
			expectedErr: nil,
		},
		{
			client: net.NewClient(
				"somehost",
				"somebase",
				"sometoken",
				"http",
			),
			entrypoint:  "hello",
			expected:    "http://somehost/somebasesometoken/hello",
			expectedErr: nil,
		},
		{
			client: net.NewClient(
				"",
				"base",
				"token",
				"protocol",
			),
			entrypoint:  "entrypoint",
			expected:    "",
			expectedErr: net.ErrNoHostProvided,
		},
		{
			client: net.NewClient(
				"host",
				"",
				"token",
				"protocol",
			),
			entrypoint:  "entrypoint",
			expected:    "",
			expectedErr: net.ErrNoBaseProvided,
		},
		{
			client: net.NewClient(
				"host",
				"base",
				"",
				"protocol",
			),
			entrypoint:  "entrypoint",
			expected:    "",
			expectedErr: net.ErrNoTokenProvided,
		},
		{
			client: net.NewClient(
				"host",
				"base",
				"token",
				"",
			),
			entrypoint:  "entrypoint",
			expected:    "",
			expectedErr: net.ErrNoProtocolProvided,
		},
		{
			client: net.NewClient(
				"host",
				"base",
				"token",
				"protocol",
			),
			entrypoint:  "",
			expected:    "protocol://host/basetoken",
			expectedErr: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.entrypoint, func(t *testing.T) {
			url, err := c.client.BuildURL(c.entrypoint)

			if err != nil {
				if errors.Is(err, c.expectedErr) {
					return
				}

				t.Errorf(`unexpected error: wanted "%s", got "%s"`, c.expectedErr, err)
			}

			if url != c.expected {
				t.Errorf("unexpected URL: got %v want %v", url, c.expected)
			}
		})
	}

}
