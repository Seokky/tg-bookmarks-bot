package endpoints

import (
	"tg-bookmarks-bot/internal/domain/entities"
)

// GetUpdatesResponse describes structure of /getUpdates response
type GetUpdatesResponse struct {
	Ok     bool              `json:"ok"`
	Result []entities.Update `json:"result"`
}

const GetUpdatesEndpoint = "/getUpdates"
