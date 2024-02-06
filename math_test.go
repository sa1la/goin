package goin

import (
	"fmt"
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
func ExampleChMax() {
	a := 1
	ChMax(&a, 2)
	fmt.Println(a)
	// Output: 2
}

func ExampleChMin() {
	a := 3
	ChMin(&a, 2)
	fmt.Println(a)
	// Output: 2
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

func TestCombo(t *testing.T) {
	tests := []struct {
		n      int
		k      int
		expect int
	}{
		{5, 3, 10},
		{4, 2, 6},
		{6, 0, 1},
		{6, 1, 6},
		{6, 6, 1},
	}

	for _, test := range tests {
		result := Combo(test.n, test.k)
		if result != test.expect {
			t.Errorf("Combo(%d, %d) = %d; expect %d", test.n, test.k, result, test.expect)
		}
	}
}

func ExampleCombo() {
	fmt.Println(Combo(5, 3))
	fmt.Println(Combo(4, 2))
	fmt.Println(Combo(6, 0))
	fmt.Println(Combo(6, 1))
	fmt.Println(Combo(6, 6))
	// Output:
	// 10
	// 6
	// 1
	// 6
	// 1
}
