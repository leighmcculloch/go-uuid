package uuidv4_test

import (
	"4d63.com/uuid/uuidv4"
)

func ExampleNew() {
	_ = uuidv4.New()
	// Returns a UUID. e.g. 08827178-0ad4-11e7-b5df-b3f54921aa61
}
