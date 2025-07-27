package entities

type Message struct {
	ID   uint64 `json:"message_id"`
	From User   `json:"from"`
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

type Chat struct {
	ID uint64 `json:"id"`
}
