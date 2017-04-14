package uuid

import (
	"encoding/hex"
	"fmt"
)

// UUID is an implementation of the UUID RFC4122 standard.
// Ref: https://www.ietf.org/rfc/rfc4122.txt
type UUID [16]byte

// New creates a UUID from a [16]byte.
func New(b [16]byte) UUID {
	return UUID(b)
}

// Parse parses the hyphenated UUID string for a UUID.
func Parse(s string) (UUID, error) {
	var id UUID
	_, err := hex.Decode(id[0:4], []byte(s[0:8]))
	if err != nil {
		return UUID{}, fmt.Errorf("%s: invalid format, want: 32 hex, with or without hyphens, %s", s, err)
	}
	_, err = hex.Decode(id[4:6], []byte(s[9:13]))
	if err != nil {
		return UUID{}, fmt.Errorf("%s: invalid format, want: 32 hex, with or without hyphens, %s", s, err)
	}
	_, err = hex.Decode(id[6:8], []byte(s[14:18]))
	if err != nil {
		return UUID{}, fmt.Errorf("%s: invalid format, want: 32 hex, with or without hyphens, %s", s, err)
	}
	_, err = hex.Decode(id[8:10], []byte(s[19:23]))
	if err != nil {
		return UUID{}, fmt.Errorf("%s: invalid format, want: 32 hex, with or without hyphens, %s", s, err)
	}
	_, err = hex.Decode(id[10:16], []byte(s[24:36]))
	if err != nil {
		return UUID{}, fmt.Errorf("%s: invalid format, want: 32 hex, with or without hyphens, %s", s, err)
	}
	return id, nil
}

// String returns a text hyphenated hexadecimal representation of the UUID.
func (id UUID) String() string {
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

// Timestamp returns the integer value of the time portion of the UUID in 100ns
// intervals since Oct 15, 1582.
//
// Note: All UUIDs have a time bits, but only TimeUUIDs (UUID v1) contain the
// time in these bits, making the return value of this function relatively
// useless for other UUIDs.
func (id UUID) Timestamp() int64 {
	return int64(id[0])<<24 |
		int64(id[1])<<16 |
		int64(id[2])<<8 |
		int64(id[3]) |
		int64(id[4])<<40 |
		int64(id[5])<<32 |
		int64(id[6]&0xdf)<<56 | // 0xdf removes version
		int64(id[7])<<48
}
