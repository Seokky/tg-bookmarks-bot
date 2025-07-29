package entities

// Message struct describes message received from polling Telegram server
type Message struct {
	ID   uint64 `json:"message_id"`
	From User   `json:"from"`
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

// Chat struct describes chat received from polling Telegram server
type Chat struct {
	ID uint64 `json:"id"`
}
