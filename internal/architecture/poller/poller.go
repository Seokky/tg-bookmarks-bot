// Package poller
// Polls the server and returns the updates to the operator
package poller

import (
	"fmt"
	"tg-bookmarks-bot/internal/domain/endpoints"
	"tg-bookmarks-bot/internal/domain/entities"
	"tg-bookmarks-bot/internal/infrastructure/net"
	"time"
)

func Start(updates chan entities.Update) {
	for {
		response, err := poll[endpoints.GetUpdatesResponse]()

		if (err == nil) && response.Ok {
			for _, update := range response.Result {
				updates <- update
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func poll[R any]() (R, error) {
	fmt.Println("poll()")

	res, err := net.Fetch[R](endpoints.GetUpdatesEndpoint)
	if err != nil {
		return res, err
	}

	return res, nil
}
