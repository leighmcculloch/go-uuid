package uuid

import (
	"testing"
)

func TestString(t *testing.T) {
	id := UUID{0x08, 0xc5, 0x1b, 0xc6, 0x01, 0xe7, 0x0a, 0x13, 0xbf, 0x3c, 0x3e, 0xa7, 0x83, 0xce, 0xe1, 0xc2}
	expectedString := "08c51bc6-01e7-0a13-bf3c-3ea783cee1c2"
	if s := id.String(); s != expectedString {
		t.Fatal([]byte(id[:]), "String() returned", s, "want", expectedString)
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		String string
		ID     UUID
	}{
		{"0dca7183-01e7-0a15-0d75-148a1db66d05", UUID{0x0d, 0xca, 0x71, 0x83, 0x01, 0xe7, 0x0a, 0x15, 0x0d, 0x75, 0x14, 0x8a, 0x1d, 0xb6, 0x6d, 0x05}},
		{"0dca718301e70a150d75148a1db66d05", UUID{0x0d, 0xca, 0x71, 0x83, 0x01, 0xe7, 0x0a, 0x15, 0x0d, 0x75, 0x14, 0x8a, 0x1d, 0xb6, 0x6d, 0x05}},
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
