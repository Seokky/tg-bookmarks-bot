package endpoints

import "tg-bookmarks-bot/internal/domain/entities"

// SendMessageParams describes schema of /sendMessage request body
type SendMessageParams struct {
	ChatID uint64 `json:"chat_id"`
	Text   string `json:"text"`
}

// SendMessageResponse describes schema of /sendMessage response body
type SendMessageResponse struct {
	Ok      bool             `json:"ok"`
	Message entities.Message `json:"message"`
}

// SendMessageEndpoint describes path of endpoint to send message to user
const SendMessageEndpoint = "/sendMessage"
