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

// Lcm calculates the least common multiple of two integers
func Lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return Abs(a*b) / Gcd(a, b)
}

// IsPrime checks if a number is prime
// Time complexity: O(sqrt(n))
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Factorial calculates n! (n factorial)
func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// ModPow calculates (base^exp) % mod using fast exponentiation
// Time complexity: O(log exp)
func ModPow(base, exp, mod int) int {
	if mod == 1 {
		return 0
	}
	result := 1
	base = base % mod
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		exp = exp >> 1
		base = (base * base) % mod
	}
	return result
}

// ModInv calculates the modular multiplicative inverse of a modulo m
// Returns -1 if inverse doesn't exist
func ModInv(a, m int) int {
	if Gcd(a, m) != 1 {
		return -1 // Inverse doesn't exist
	}
	return ModPow(a, m-2, m) // Works when m is prime
}

// ExtendedGcd calculates gcd(a, b) and finds x, y such that ax + by = gcd(a, b)
func ExtendedGcd(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x1, y1 := ExtendedGcd(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return gcd, x, y
}

// SieveOfEratosthenes generates all prime numbers up to n
func SieveOfEratosthenes(n int) []int {
	if n < 2 {
		return []int{}
	}
	
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	
	var primes []int
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

// Fibonacci calculates the nth Fibonacci number
func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// IsPowerOfTwo checks if n is a power of 2
func IsPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

// NextPowerOfTwo returns the smallest power of 2 that is >= n
func NextPowerOfTwo(n int) int {
	if n <= 1 {
		return 1
	}
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	return n + 1
}
