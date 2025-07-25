// Package poller
// polls the Telegram server and writes updates to updates channel
package poller

import (
	"fmt"
	"strconv"
	"tg-bookmarks-bot/internal/domain/endpoints"
	"tg-bookmarks-bot/internal/domain/entities"
	"tg-bookmarks-bot/internal/infrastructure/net"
	"time"
)

const (
	errorSleepTimeout = 3  // Delay in seconds after poll request failed
	longPollTimeout   = 10 // Timeout in seconds to long polling waiting
)

// Start starts polling Telegram server to get updates
func Start(updates chan<- entities.Update) {
	var lastUpdateID uint64

	for {
		response, err := poll(lastUpdateID)

		if err != nil {
			time.Sleep(errorSleepTimeout * time.Second)
			continue
		}

		if response.Ok {
			for _, update := range response.Result {
				updates <- update

				lastUpdateID = update.ID
			}
		}
	}
}

// Make long polling request
func poll(fromID uint64) (endpoints.GetUpdatesResponse, error) {
	params := longPollingParams(fromID)

	res, err := net.Fetch[endpoints.GetUpdatesResponse](endpoints.GetUpdatesEndpoint, params)

	if err != nil {
		return res, err
	}

	return res, nil
}

// Prepare and return query params for long polling request
func longPollingParams(fromID uint64) net.FetchQueryParams {
	params := net.FetchQueryParams{
		{
			Name:  "timeout",
			Value: strconv.Itoa(longPollTimeout),
		},
	}

	if fromID != 0 {
		params = append(params, net.FetchQueryParam{
			Name:  "offset",
			Value: fmt.Sprint(fromID + 1),
		})
	}

	return params
}
