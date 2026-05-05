package goin

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Max 返回一组可比较值中的最大值。调用方需保证至少传入一个值。
func Max[T constraints.Ordered](as ...T) T {
	res := as[0]
	for _, a := range as {
		if res < a {
			res = a
		}
	}
	return res
}

// Min 返回一组可比较值中的最小值。调用方需保证至少传入一个值。
func Min[T constraints.Ordered](as ...T) T {
	res := as[0]
	for _, a := range as {
		if res > a {
			res = a
		}
	}
	return res
}

// ChMax 若 b > *a，则将 *a 更新为 b。
func ChMax(a *int, b int) {
	*a = Max(*a, b)
}

// ChMin 若 b < *a，则将 *a 更新为 b。
func ChMin(a *int, b int) {
	*a = Min(*a, b)
}

// Abs 返回整数 a 的绝对值。
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Pow 计算 x 的 n 次幂（快速幂），时间复杂度 O(log n)。
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

// Sum 返回一组整数的和。
func Sum(s ...int) int {
	res := 0
	for _, v := range s {
		res += v
	}
	return res
}

// GetAngle 计算从正 x 轴到点 (x, y) 的角度（单位：度）。
func GetAngle(x, y float64) float64 {
	return math.Atan2(y, x) * 180 / math.Pi
}

// Combo 计算组合数 C(n, k)，即 n 选 k 的方案数。
// 使用递推公式 C(n, k) = C(n, k-1) * (n-k+1) / k。
func Combo(n int, k int) int {
	if k == 0 {
		return 1
	}
	return (n - k + 1) * Combo(n, k-1) / k
}

// Gcd 返回 a 和 b 的最大公约数。
func Gcd(a, b int) int {
	if b == 0 {
		return Abs(a)
	}
	return Gcd(b, a%b)
}

// Lcm 返回 a 和 b 的最小公倍数。
// 若 a 或 b 为 0，返回 0。
func Lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return Abs(a*b) / Gcd(a, b)
}

// IsPrime 判断 n 是否为质数，时间复杂度 O(sqrt(n))。
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

// Factorial 返回 n 的阶乘 n!。
// n < 0 时返回 0。
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

// ModPow 使用快速幂计算 (base^exp) mod mod，时间复杂度 O(log exp)。
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

// ModInv 返回 a 在模 m 下的乘法逆元。
// 若逆元不存在（即 gcd(a, m) != 1），返回 -1。
// 当 m 为质数时，内部使用费马小定理（ModPow）计算。
func ModInv(a, m int) int {
	if Gcd(a, m) != 1 {
		return -1
	}
	return ModPow(a, m-2, m)
}

// ExtendedGcd 返回 gcd(a, b)，并求出满足 ax + by = gcd(a, b) 的 x 和 y。
func ExtendedGcd(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x1, y1 := ExtendedGcd(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return gcd, x, y
}

// SieveOfEratosthenes 返回不超过 n 的所有质数（埃氏筛），时间复杂度 O(n log log n)。
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

// Fibonacci 返回第 n 个斐波那契数。
// F(0)=0, F(1)=1。
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

// IsPowerOfTwo 判断 n 是否为 2 的幂。
func IsPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

// NextPowerOfTwo 返回不小于 n 的最小 2 的幂。
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
