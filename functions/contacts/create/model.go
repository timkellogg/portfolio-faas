package main

import (
	"github.com/lib/pq"
)

type contact struct {
	ID        string      `json:"id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Email     string      `json:"email"`
	Phone     string      `json:"phone"`
	Message   string      `json:"message"`
	CreatedAt pq.NullTime `json:"created_at"`
}
