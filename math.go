package goin

import (
	"math"

	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](as ...T) T {
	res := as[0]
	for _, a := range as {
		if res < a {
			res = a
		}
	}
	return res
}
func Min[T constraints.Ordered](as ...T) T {
	res := as[0]
	for _, a := range as {
		if res > a {
			res = a
		}
	}
	return res
}
func ChMax(a *int, b int) {
	*a = Max(*a, b)
}
func ChMin(a *int, b int) {
	*a = Min(*a, b)
}
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// The Pow function is used to calculate the power of x to the n
// x is the base, n is the exponent
// It uses the method of fast power to calculate, with a time complexity of O(logn)
// If n is even, then x^n = (x^2)^(n/2)
// If n is odd, then x^n = x * x^(n-1)
// By continuously halving, the exponent n can be reduced to 0, at which point x^0 = 1
// In the process of halving, if n is odd, an extra x needs to be multiplied
func Pow(x, n int) int {
	res := 1
	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		x *= x
		n >>= 1
	}
	return res
}

func Sum(s ...int) int {
	res := 0
	for _, v := range s {
		res += v
	}
	return res
}

// The GetAngle function calculates the angle (in degrees) between the positive x-axis and the point given by the coordinates (x, y).
// It uses the Atan2 function from the math package to compute the arc tangent of y/x in radians and then converts it to degrees.
func GetAngle(x, y float64) float64 {
	return math.Atan2(y, x) * 180 / math.Pi
}

// The Combo function is used to calculate the combination number C(n, k)
// n is the total number, k is the number of choices
// If k is 0, return 1, indicating that there is only one situation without choice
// Otherwise, the Combo function is called recursively, reducing the value of k each time until k is 0
// Each recursion will multiply by (n - k + 1) / k, which is the calculation formula of the combination number
func Combo(n int, k int) int {
	if k == 0 {
		return 1
	}
	return (n - k + 1) * Combo(n, k-1) / k
}

func Gcd(a, b int) int {
	if b == 0 {
		return Abs(a)
	}
	return Gcd(b, a%b)
}
