package responser

import (
	"tg-bookmarks-bot/internal/domain/endpoints"
	"tg-bookmarks-bot/internal/infrastructure/net"
)

func Send(chatID uint64, text string) {
	params := map[string]any{
		"chat_id": chatID,
		"text":    text,
	}

	_, _ = net.Send[endpoints.SendMessageResponse](endpoints.SendMessageEndpoint, params)
}
