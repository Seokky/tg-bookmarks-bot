package entities

type Message struct {
	ID   int  `json:"message_id"`
	From User `json:"from"`
}
