package mangle

import (
	"testing"
)

func BenchmarkDot1(b *testing.B) {
	x := Vector4d{1, 1.23, 1, 3.14}
	y := Vector4d{2, 2, 2, 2}
	for i := 0; i < b.N; i++ {
		Dot(&x, &y)
	}
}

func TestCap1(t *testing.T) {
	cap := Cap{Vector4d{1, 0, 0, -1}, 0.5}
	v1 := Vector4d{1, 0, 0, 1}
	if !InCap(&cap, &v1) {
		t.Error("Failed testing v1")
	}
	v2 := Vector4d{1, 1, 0, 0}
	if InCap(&cap, &v2) {
		t.Error("Failed testing v2")
	}
}

func BenchmarkCap1(b *testing.B) {
	cap := Cap{Vector4d{1, 0, 0, -1}, 0.5}
	v1 := Vector4d{1, 0, 0, 1}
	v2 := Vector4d{1, 1, 0, 0}
	for i := 0; i < b.N; i++ {
		InCap(&cap, &v1)
		InCap(&cap, &v2)
	}
}
