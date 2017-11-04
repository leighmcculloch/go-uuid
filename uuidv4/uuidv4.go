package uuidv4

import (
	"crypto/rand"

	"4d63.com/uuid"
)

// New creates a random UUID.
func New() uuid.UUID {
	u := uuid.UUID{}
	rand.Read(u[:])
	return u
}
