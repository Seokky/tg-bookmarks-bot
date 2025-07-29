package entities

// User struct describes user received from polling Telegram server
type User struct {
	ID   int    `json:"id"`
	Name string `json:"username"`
}
