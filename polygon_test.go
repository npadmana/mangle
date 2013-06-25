package mangle

import (
	"math"
	"testing"
)

func NotNear(x, y, eps float64) bool {
	return eps < math.Abs(x-y)
}

func TestParsePolygon1(t *testing.T) {
	p := &Polygon{}
	p.Parse([]byte("polygon       2753 ( 4 caps, 0.333333333333333 weight, 1904 pixel, 3.3146470414e-05 str):"))
	if p.id != 2753 {
		t.Error("Wrong id")
	}
	if p.ncaps != 4 {
		t.Error("wrong ncaps ", p.ncaps)
	}
	if len(p.clist) != 4 {
		t.Error("wrong caplist")
	}
	if NotNear(p.weight, 0.333333333333333, 1.e-10) {
		t.Error("wrong weight")
	}

}

func TestParsePolygon2(t *testing.T) {
	p := &Polygon{}
	p.Parse([]byte("polygon       3367 ( 4 caps, 0.970588235294118 weight, 1968 pixel, 2.9056936228e-05 str)"))
	if p.id != 3367 {
		t.Error("Wrong id")
	}
	if p.ncaps != 4 {
		t.Error("wrong ncaps ", p.ncaps)
	}
	if len(p.clist) != 4 {
		t.Error("wrong caplist")
	}
	if NotNear(p.weight, 0.970588235294118, 1.e-10) {
		t.Error("wrong weight")
	}

}
