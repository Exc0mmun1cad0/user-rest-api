package entity

import (
	"time"
)

type User struct {
	ID        string    `json:"id" db:"user_id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	Age       uint      `json:"age" db:"age"`
	Created   time.Time `json:"created_at" db:"created_at"`
}
