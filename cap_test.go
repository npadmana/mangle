package mangle

import (
	"strings"
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
	cap := Cap{vector4d{1, 0, 0, -1}, 0.5}
	v1 := vector4d{1, 0, 0, 1}
	if !cap.In(&v1) {
		t.Error("Failed testing v1")
	}
	v2 := vector4d{1, 1, 0, 0}
	if cap.In(&v2) {
		t.Error("Failed testing v2")
	}
}

func TestCap2(t *testing.T) {
	c := Cap{}
	if err := c.Read(strings.NewReader("0 0 1 0.5\n")); err != nil {
		t.Error(err)
	}
	v1 := vector4d{1, 0, 0, 1}
	if !c.In(&v1) {
		t.Error("Failed testing v1")
	}
	v2 := vector4d{1, 1, 0, 0}
	if c.In(&v2) {
		t.Error("Failed testing v2")
	}
}

func TestCap3(t *testing.T) {
	c := &Cap{}
	if err := c.Read(strings.NewReader("0 0 1 -0.5\n")); err != nil {
		t.Error(err)
	}
	v1 := vector4d{1, 0, 0, 1}
	if c.In(&v1) {
		t.Error("Failed testing v1")
	}
	v2 := vector4d{1, 1, 0, 0}
	if !c.In(&v2) {
		t.Error("Failed testing v2")
	}
}

func BenchmarkCap1(b *testing.B) {
	cap := Cap{vector4d{1, 0, 0, -1}, 0.5}
	v1 := vector4d{1, 0, 0, 1}
	v2 := vector4d{1, 1, 0, 0}
	for i := 0; i < b.N; i++ {
		cap.In(&v1)
		cap.In(&v2)
	}
}
