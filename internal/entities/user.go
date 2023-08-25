package entities

// User ...
type User struct {
	ID      uint   `json:"-"`
	Name    string `json:"-"`
	Address string `json:"-"`
}
