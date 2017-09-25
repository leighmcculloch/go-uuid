package timeuuid

import (
	"crypto/rand"
	"time"

	"4d63.com/uuid"
)

var nodeID [6]byte
var clockSeq [2]byte
var version byte = 0x1

func init() {
	rand.Read(nodeID[:])
	rand.Read(clockSeq[:])
}

// Now creates a UUID using the now time.
func Now() uuid.UUID {
	return New(time.Now())
}

// New creates a UUID using a given time.
func New(t time.Time) uuid.UUID {
	time := getTimestampFromTime(t)
	return uuid.UUID{
		// 0-3: 4 bytes time low
		byte(time >> 24),
		byte(time >> 16),
		byte(time >> 8),
		byte(time),
		// 4-5: 2 bytes time mid
		byte(time >> 40),
		byte(time >> 32),
		// 6-7: 2 bytes time hi or'd with version
		byte(time>>56) | (version << 4),
		byte(time >> 48),
		// 8-9: 2 bytes clock seq
		clockSeq[0] | 0x01,
		clockSeq[1],
		// 10-16: 6 bytes node ID
		nodeID[0] | 0x01, // set the multicast bit, to indicate this is not a network MAC
		nodeID[1],
		nodeID[2],
		nodeID[3],
		nodeID[4],
		nodeID[5],
	}
}
