package endpoints

import (
	"tg-bookmarks-bot/internal/domain/entities"
)

// GetUpdatesResponse describes structure of /getUpdates response
type GetUpdatesResponse struct {
	Ok     bool              `json:"ok"`
	Result []entities.Update `json:"result"`
}

// GetUpdatesEndpoint describes path of endpoint to get updates by polling Telegram server
const GetUpdatesEndpoint = "/getUpdates"
