package uuidv1

import (
	"time"
)

// Ref: https://support.datastax.com/hc/en-us/articles/204226019-Converting-TimeUUID-Strings-to-Dates
const timeOffset int64 = 122192928000000000

func getTimestampFromTime(t time.Time) int64 {
	return timeOffset + t.UTC().UnixNano()/100
}

func getTimeFromTimestamp(timestamp int64) time.Time {
	return time.Unix(0, (timestamp-timeOffset)*100)
}
