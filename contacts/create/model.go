package main

import (
	"github.com/lib/pq"
)

// Contact - client message
type Contact struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Message   string
	CreatedAt pq.NullTime
}
