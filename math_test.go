package goin

import "testing"

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
}

func TestSum(t *testing.T) {
	if Sum(1, 2, 3) != 6 {
		t.Fail()
	}
}
