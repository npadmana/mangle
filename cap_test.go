package mangle

import (
	"testing"
)

func BenchmarkDot1(b *testing.B) {
	x := vector4d{1, 1.23, 1, 3.14}
	y := vector4d{2, 2, 2, 2}
	for i := 0; i < b.N; i++ {
		dot(&x, &y)
	}
}

func TestCap1(t *testing.T) {
	cap := cap{vector4d{1, 0, 0, -1}, 0.5}
	v1 := vector4d{1, 0, 0, 1}
	if !incap(&cap, &v1) {
		t.Error("Failed testing v1")
	}
	v2 := vector4d{1, 1, 0, 0}
	if incap(&cap, &v2) {
		t.Error("Failed testing v2")
	}
}

func BenchmarkCap1(b *testing.B) {
	cap := cap{vector4d{1, 0, 0, -1}, 0.5}
	v1 := vector4d{1, 0, 0, 1}
	v2 := vector4d{1, 1, 0, 0}
	for i := 0; i < b.N; i++ {
		incap(&cap, &v1)
		incap(&cap, &v2)
	}
}
