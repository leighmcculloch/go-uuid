package timeuuid

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	t.Log(Now())
}

func TestNowUnique(t *testing.T) {
	var seen map[TimeUUID]bool

	for i := 0; i < 1000; i++ {
		id := Now()
		if _, ok := seen[id]; ok {
			t.Fatal("Seen", id, "already, expected unique")
		}
	}
}

func TestNowOrder(t *testing.T) {
	var lastID TimeUUID
	for i := 0; i < 100000; i++ {
		id := Now()
		if id.String() < lastID.String() {
			t.Fatal(id, "sorts earlier than last", lastID, "want to sort later")
		}
		lastID = id
	}
}

func TestNewOrder(t *testing.T) {
	testPeriodYears := 20
	timeStart := time.Now().AddDate(-testPeriodYears/2, 0, 0)
	timeEnd := timeStart.AddDate(testPeriodYears, -1, 0)
	var lastID TimeUUID
	for time := timeStart; time.Before(timeEnd); time = time.AddDate(0, 0, 1) {
		id := New(time)
		if id.Timestamp() < lastID.Timestamp() {
			t.Fatal(id, "sorts earlier than last", lastID, "want to sort later")
		}
		lastID = id
	}
}

func TestString(t *testing.T) {
	id := TimeUUID{0x08, 0xc5, 0x1b, 0xc6, 0x01, 0xe7, 0x0a, 0x13, 0xbf, 0x3c, 0x3e, 0xa7, 0x83, 0xce, 0xe1, 0xc2}
	expectedString := "08c51bc6-01e7-0a13-bf3c-3ea783cee1c2"
	if s := id.String(); s != expectedString {
		t.Fatal([]byte(id[:]), "String() returned", s, "want", expectedString)
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		String string
		ID     TimeUUID
	}{
		{"0dca7183-01e7-0a15-0d75-148a1db66d05", TimeUUID{0x0d, 0xca, 0x71, 0x83, 0x01, 0xe7, 0x0a, 0x15, 0x0d, 0x75, 0x14, 0x8a, 0x1d, 0xb6, 0x6d, 0x05}},
		{"0dca718301e70a150d75148a1db66d05", TimeUUID{0x0d, 0xca, 0x71, 0x83, 0x01, 0xe7, 0x0a, 0x15, 0x0d, 0x75, 0x14, 0x8a, 0x1d, 0xb6, 0x6d, 0x05}},
	}

	for _, test := range tests {
		id, err := Parse(test.String)
		if err != nil {
			t.Errorf("String %s: got error %s", test.String, err)
		} else if id != test.ID {
			t.Errorf("String %s: got %v, want %v", test.String, id, test.ID)
		}
	}
}

func TestParseError(t *testing.T) {
	tests := []struct {
		String string
		Error  string
	}{
		{"0dca7183*01e7*0a15*0d75*148a1db66d05", "0dca7183*01e7*0a15*0d75*148a1db66d05: invalid format, want: 32 hex, with or without hyphens"},
		{"helloworld", "helloworld: invalid format, want: 32 hex, with or without hyphens"},
	}

	for _, test := range tests {
		_, err := Parse(test.String)
		if err == nil {
			t.Errorf("String %s: got no error, want %s", test.String, test.Error)
		} else if err.Error() != test.Error {
			t.Errorf("String %s: got error %v, want %v", test.String, err, test.Error)
		}
	}
}
