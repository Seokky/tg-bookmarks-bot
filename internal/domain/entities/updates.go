package entities

type Update struct {
	ID      uint64  `json:"update_id"`
	Message Message `json:"message"`
}
