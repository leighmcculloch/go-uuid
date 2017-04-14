package uuid_test

import (
	"fmt"

	"github.com/leighmcculloch/go-uuid"
)

func Example() {
	id, _ := uuid.Parse("08827178-0ad4-11e7-b5df-b3f54921aa61")

	fmt.Println("String:", id.String())
	fmt.Println("Timestamp:", id.Timestamp())

	// Output:
	// String: 08827178-0ad4-11e7-b5df-b3f54921aa61
	// Timestamp: 1290011724057047416
}
