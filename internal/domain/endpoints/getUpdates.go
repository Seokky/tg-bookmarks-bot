package endpoints

import (
	e "tg-bookmarks-bot/internal/domain/entities"
)

// GetUpdatesResponse describes structure of /getUpdates response
type GetUpdatesResponse struct {
	Ok     bool       `json:"ok"`
	Result []e.Update `json:"result"`
}

const GetUpdatesEndpoint = "/getUpdates"
