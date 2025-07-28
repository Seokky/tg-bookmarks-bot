package endpoints

import "tg-bookmarks-bot/internal/domain/entities"

type SendMessageParams struct {
	ChatID uint64 `json:"chat_id"`
	Text   string `json:"text"`
}
type SendMessageResponse struct {
	Ok      bool             `json:"ok"`
	Message entities.Message `json:"message"`
}

const SendMessageEndpoint = "/sendMessage"
