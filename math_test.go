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
func TestGcd(t *testing.T) {
	testCases := []struct {
		a        int
		b        int
		expected int
	}{
		{10, 5, 5},
		{14, 28, 14},
		{18, 35, 1},
		{40, 100, 20},
		{-5, -15, 5},
		{0, 0, 0},
	}

	for _, tc := range testCases {
		result := Gcd(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Gcd(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
		}
	}
}

func TestLcm(t *testing.T) {
	testCases := []struct {
		a        int
		b        int
		expected int
	}{
		{4, 6, 12},
		{15, 25, 75},
		{7, 5, 35},
		{0, 5, 0},
		{12, 18, 36},
	}

	for _, tc := range testCases {
		result := Lcm(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Lcm(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
		}
	}
}

func TestIsPrime(t *testing.T) {
	testCases := []struct {
		n        int
		expected bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{17, true},
		{25, false},
		{29, true},
		{1, false},
		{0, false},
		{-5, false},
	}

	for _, tc := range testCases {
		result := IsPrime(tc.n)
		if result != tc.expected {
			t.Errorf("IsPrime(%d) = %v; expected %v", tc.n, result, tc.expected)
		}
	}
}

func TestFactorial(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
		{7, 5040},
		{-1, 0},
	}

	for _, tc := range testCases {
		result := Factorial(tc.n)
		if result != tc.expected {
			t.Errorf("Factorial(%d) = %d; expected %d", tc.n, result, tc.expected)
		}
	}
}

func TestModPow(t *testing.T) {
	testCases := []struct {
		base     int
		exp      int
		mod      int
		expected int
	}{
		{2, 10, 1000, 24},
		{3, 4, 5, 1},
		{5, 3, 13, 8},
		{2, 0, 5, 1},
	}

	for _, tc := range testCases {
		result := ModPow(tc.base, tc.exp, tc.mod)
		if result != tc.expected {
			t.Errorf("ModPow(%d, %d, %d) = %d; expected %d", tc.base, tc.exp, tc.mod, result, tc.expected)
		}
	}
}

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{10, 55},
		{15, 610},
	}

	for _, tc := range testCases {
		result := Fibonacci(tc.n)
		if result != tc.expected {
			t.Errorf("Fibonacci(%d) = %d; expected %d", tc.n, result, tc.expected)
		}
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	testCases := []struct {
		n        int
		expected bool
	}{
		{1, true},
		{2, true},
		{4, true},
		{8, true},
		{16, true},
		{3, false},
		{5, false},
		{6, false},
		{0, false},
		{-4, false},
	}

	for _, tc := range testCases {
		result := IsPowerOfTwo(tc.n)
		if result != tc.expected {
			t.Errorf("IsPowerOfTwo(%d) = %v; expected %v", tc.n, result, tc.expected)
		}
	}
}

func TestSieveOfEratosthenes(t *testing.T) {
	result := SieveOfEratosthenes(30)
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	
	if len(result) != len(expected) {
		t.Errorf("SieveOfEratosthenes(30) length = %d; expected %d", len(result), len(expected))
		return
	}
	
	for i, prime := range expected {
		if result[i] != prime {
			t.Errorf("SieveOfEratosthenes(30)[%d] = %d; expected %d", i, result[i], prime)
		}
	}
}
