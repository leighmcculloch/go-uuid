package uuid_test

import (
	"github.com/leighmcculloch/go-uuid/timeuuid"
)

func Example_timeUUIDNow() {
	_ = timeuuid.Now()
	// Returns a UUID. e.g. 08827178-0ad4-11e7-b5df-b3f54921aa61
}
