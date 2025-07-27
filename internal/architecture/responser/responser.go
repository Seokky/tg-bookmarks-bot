package responser

import (
	"tg-bookmarks-bot/internal/domain/endpoints"
	"tg-bookmarks-bot/internal/infrastructure/net"
)

func Send(chatId uint64, text string) {
	params := map[string]any{
		"chat_id": chatId,
		"text":    text,
	}

	_, _ = net.Send[endpoints.SendMessageResponse](endpoints.SendMessageEndpoint, params)
}
