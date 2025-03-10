package main

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Age       uint      `json:"age"`
	Created   time.Time `json:"created_at"`
}

// I know it's awful. I'll fix it in next versions
var (
	db = make(map[string]User, 0)
	mu sync.Mutex
)
