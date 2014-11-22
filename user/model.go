package user

import "time"

type User struct {
	id       int
	email    string
	password string

	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}
