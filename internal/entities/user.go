package entities

import "time"

// User ...
type User struct {
	ID             uint      `json:"-"`
	Name           string    `json:"-"`
	Surname        string    `json:"-"`
	DateOfBirthday time.Time `json:"-"`
}
