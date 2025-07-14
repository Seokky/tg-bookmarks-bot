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

func Start(updates chan<- entities.Update) {
	fmt.Println("poller started")

	for {
		poll()

		time.Sleep(1 * time.Second)
	}
}

func poll() {
	fmt.Println("poll()")
	res, _ := net.Fetch[endpoints.GetUpdatesResponse](endpoints.GetUpdatesEndpoint)
	fmt.Printf("%+v\n", res)
}
