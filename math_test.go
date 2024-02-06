package goin

import (
	"math"
	"testing"
)

func TestMax(t *testing.T) {
	if Max(1, 2, 3) != 3 {
		t.Fail()
	}
}

func TestMin(t *testing.T) {
	if Min(1, 2, 3) != 1 {
		t.Fail()
	}
}

func TestChMax(t *testing.T) {
	a := 1
	ChMax(&a, 2)
	if a != 2 {
		t.Fail()
	}
}

func TestChMin(t *testing.T) {
	a := 3
	ChMin(&a, 2)
	if a != 2 {
		t.Fail()
	}
}

func TestAbs(t *testing.T) {
	if Abs(-1) != 1 {
		t.Fail()
	}
}

func TestPow(t *testing.T) {
	if Pow(2, 3) != 8 {
		t.Fail()
	}
	if Pow(3, 2) != 9 {
		t.Fail()
	}
	if Pow(2, 0) != 1 {
		t.Fail()
	}
	if Pow(0, 2) != 0 {
		t.Fail()
	}
}

func TestSum(t *testing.T) {
	if Sum(1, 2, 3) != 6 {
		t.Fail()
	}
}

func TestGetAngle(t *testing.T) {
	testCases := []struct {
		x        float64
		y        float64
		expected float64
	}{
		{1, 0, 0},
		{0, 1, 90},
		{-1, 0, 180},
		{0, -1, -90},
		{1, 1, 45},
	}

	for _, tc := range testCases {
		result := GetAngle(tc.x, tc.y)
		if math.Abs(result-tc.expected) > 1e-9 {
			t.Errorf("getAngle(%v, %v) = %v, expected %v", tc.x, tc.y, result, tc.expected)
		}
	}
}
