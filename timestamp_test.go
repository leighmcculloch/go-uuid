package timeuuid

import (
	"testing"
	"time"
)

var timestampValidMappings = []struct {
	Time      time.Time
	Timestamp int64
}{
	{time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC), 122192928000000000},
	{time.Date(2012, time.November, 14, 20, 16, 22, 0, time.UTC), 135722169820000000},
}

func TestGetTimestampFromTime(t *testing.T) {
	for _, mapping := range timestampValidMappings {
		i := getTimestampFromTime(mapping.Time)
		if i != mapping.Timestamp {
			t.Error("Time", mapping.Time, "mapped to", i, ", want", mapping.Timestamp)
		}
	}
}

func TestGetTimeFromTimestamp(t *testing.T) {
	for _, mapping := range timestampValidMappings {
		time := getTimeFromTimestamp(mapping.Timestamp)
		if !time.Equal(mapping.Time) {
			t.Error("Int64", mapping.Timestamp, "mapped to", time, ", want", mapping.Time)
		}
	}
}
