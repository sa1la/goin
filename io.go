package goin

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

const Inf = math.MaxInt64
const BaseRune = 'a'

var debugFlg bool

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := os.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
		debugFlg = true
	}
}

func NextString() string {
	sc.Scan()
	return sc.Text()
}

// Next n-idx string-slice
func NextStrings(n int) []string {
	res := make([]string, n)
	for i := range res {
		res[i] = NextString()
	}
	return res
}

// Next string to rune[]
func NextRunes() []rune {
	return []rune(NextString())
}

func unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func NextInt() int {
	s := NextString()
	return unwrap(strconv.Atoi(s))
}

// Next int slice
func NextIntSlice(n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = NextInt()
	}
	return res
}

// Next n-idx int slice , with value increased start from base to base+n
func NewIncreaseIntSlice(n int, base int) []int {
	newSlice := make([]int, n)
	for i := 0; i < n; i++ {
		newSlice[i] = i + base
	}
	return newSlice
}

// Next double n-idx int slice a & b,
// input  1, 2 will be a[1] b[2]
func Next2IntSlice(n int) ([]int, []int) {
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = NextInt2()
	}
	return a, b
}
func NextIntSlice2D(n, m int) [][]int {
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = NextIntSlice(m)
	}
	return a
}

// create new (n,m)-idx int slice, return [][]int
func NewIntSlice2D(n, m, def int) [][]int {
	newSlice := make([][]int, n)
	for i := 0; i < n; i++ {
		newSlice[i] = make([]int, m)
		for j := 0; j < m; j++ {
			newSlice[i][j] = def
		}
	}
	return newSlice
}
func NextFloat() float64 {
	return unwrap(strconv.ParseFloat(NextString(), 64))
}
func NextFloats(n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = NextFloat()
	}
	return res
}
func NextInt2() (int, int) {
	return NextInt(), NextInt()
}
func NextInt3() (int, int, int) {
	return NextInt(), NextInt(), NextInt()
}
func NextInt4() (int, int, int, int) {
	return NextInt(), NextInt(), NextInt(), NextInt()
}

func Print(a ...any) {
	if _, err := fmt.Fprint(wr, a...); err != nil {
		panic(fmt.Errorf("print: %w", err))
	}
}

func Printf(format string, a ...any) {
	fmt.Fprintf(wr, format, a...)
}
func Println(a ...any) {
	fmt.Fprintln(wr, a...)
}
func Flush() {
	wr.Flush()
}
func Debug(v ...interface{}) {
	if !debugFlg {
		return
	}
	Println(v...)
}
