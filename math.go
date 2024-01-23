package goin

// simple math functions for ordered
type signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type integer interface {
	signed | unsigned
}

type float interface {
	~float32 | ~float64
}

type ordered interface {
	integer | float | ~string
}

func Max[T ordered](as ...T) T {
	res := as[0]
	for _, a := range as {
		if res < a {
			res = a
		}
	}
	return res
}
func Min[T ordered](as ...T) T {
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
func Pow(a, n int) int {
	res := 1
	b := a
	for n > 0 {
		if n&1 > 0 {
			res *= b
		}
		n >>= 1
		b *= b
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
