package rstr

import (
	"testing"
)

func TestGenstr(t *testing.T) {
	expectedLen := 64

	el := len(Genstr(expectedLen))
	ec := len(Genstr(expectedLen))

	if el != expectedLen {
		t.Errorf("ERROR... Expected length %d, got %d", expectedLen, el)
	}

	if ec != expectedLen {
		t.Errorf("ERROR... Expected length %d, got %d", expectedLen, ec)
	}
}

func BenchmarkGenstr8(b *testing.B) {
	customProps(8, b)
}

func BenchmarkGenstr16(b *testing.B) {
	customProps(16, b)
}

func BenchmarkGenstr32(b *testing.B) {
	customProps(32, b)
}

func BenchmarkGenstr64(b *testing.B) {
	customProps(64, b)
}

func customProps(l int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Genstr(l)
	}
}
