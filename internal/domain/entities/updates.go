package entities

// Update struct describes update received from polling Telegram server
type Update struct {
	ID      uint64  `json:"update_id"`
	Message Message `json:"message"`
}
