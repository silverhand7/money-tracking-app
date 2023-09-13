package domain

import "time"

type User struct {
	ID        int32
	Name      string
	Email     string
	Password  string
	ApiKey    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
