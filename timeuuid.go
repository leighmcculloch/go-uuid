package timeuuid

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// TimeUUID is an implementation of the UUID v1 RFC4122 standard
// using a randomly generated node ID and clock sequence.
// Ref: https://www.ietf.org/rfc/rfc4122.txt
type TimeUUID [16]byte

var nodeID [6]byte
var clockSeq [2]byte
var version byte = 0x1

func init() {
	rand.Read(nodeID[:])
	rand.Read(clockSeq[:])
}

// Now creates a TimeUUID using the now time.
func Now() TimeUUID {
	return New(time.Now())
}

// New creates a TimeUUID using a given time.
func New(t time.Time) TimeUUID {
	time := getTimestampFromTime(t)
	return TimeUUID{
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

// String returns a text hexadecimal representation of the TimeUUID.
func (id TimeUUID) String() string {
	var b [36]byte
	hex.Encode(b[0:8], id[0:4])
	b[8] = '-'
	hex.Encode(b[9:13], id[4:6])
	b[13] = '-'
	hex.Encode(b[14:18], id[6:8])
	b[18] = '-'
	hex.Encode(b[19:23], id[8:10])
	b[23] = '-'
	hex.Encode(b[24:36], id[10:16])
	return string(b[:])
}

// Timestamp returns the integer value of the time in 100ns intervals since Oct 15, 1582.
func (id TimeUUID) Timestamp() int64 {
	return int64(id[0])<<24 |
		int64(id[1])<<16 |
		int64(id[2])<<8 |
		int64(id[3]) |
		int64(id[4])<<40 |
		int64(id[5])<<32 |
		int64(id[6]&0xdf)<<56 | // 0xdf removes version
		int64(id[7])<<48
}

// Parse parses the string for a TimeUUID for the given textuable representation.
func Parse(s string) (TimeUUID, error) {
	stripped := strings.Replace(s, "-", "", -1)
	if len(stripped) != 32 {
		return TimeUUID{}, fmt.Errorf("%s: invalid format, want: 32 hex, with or without hyphens", s)
	}
	b, err := hex.DecodeString(stripped)
	if err != nil {
		return TimeUUID{}, fmt.Errorf("%s: invalid format, want: 32 hex, with or without hyphens, %s", s, err)
	}
	var id TimeUUID
	copy(id[:], b)
	return id, nil
}
