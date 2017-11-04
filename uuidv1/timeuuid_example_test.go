package uuidv1_test

import (
	"4d63.com/uuid/uuidv1"
)

func ExampleNow() {
	_ = uuidv1.Now()
	// Returns a UUID. e.g. 08827178-0ad4-11e7-b5df-b3f54921aa61
}
