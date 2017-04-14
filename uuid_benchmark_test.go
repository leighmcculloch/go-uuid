package uuid

import (
	"testing"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New([16]byte{0x08, 0x82, 0x71, 0x78, 0x0a, 0xd4, 0x11, 0xe7, 0xb5, 0xdf, 0xb3, 0xf5, 0x49, 0x21, 0xaa, 0x61})
	}
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse("08827178-0ad4-11e7-b5df-b3f54921aa61")
	}
}

func BenchmarkString(b *testing.B) {
	id := New([16]byte{0x08, 0x82, 0x71, 0x78, 0x0a, 0xd4, 0x11, 0xe7, 0xb5, 0xdf, 0xb3, 0xf5, 0x49, 0x21, 0xaa, 0x61})
	for i := 0; i < b.N; i++ {
		id.String()
	}
}