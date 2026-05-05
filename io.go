package goin

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

var (
	// sc 是全局词法扫描器，默认从标准输入按空白字符分词读取。
	// 测试中可以重新绑定到其他 Reader 以注入输入。
	sc = bufio.NewScanner(os.Stdin)
	// wr 是全局缓冲写入器，默认写入标准输出。
	// 任何输出函数调用后，程序结束前必须调用 Flush()，否则输出可能丢失。
	wr = bufio.NewWriter(os.Stdout)
)

const (
	// Inf 表示一个足够大的整数，常用于图论或动态规划中的"无穷大"。
	Inf = math.MaxInt64
	// BaseRune 是字符 'a' 的 rune 值，常用于字符与数字之间的转换。
	BaseRune = 'a'
)

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

// NextString 读取下一个以空白字符分隔的字符串。
func NextString() string {
	sc.Scan()
	return sc.Text()
}

// NextStrings 读取 n 个字符串并返回字符串切片。
func NextStrings(n int) []string {
	res := make([]string, n)
	for i := range res {
		res[i] = NextString()
	}
	return res
}

// NextRunes 读取下一个字符串并将其转为 rune 切片返回。
func NextRunes() []rune {
	return []rune(NextString())
}

// unwrap 解包 (T, error) 对；若 err != nil 则 panic。
func unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

// NextInt 读取下一个整数。若无法解析则 panic。
func NextInt() int {
	s := NextString()
	return unwrap(strconv.Atoi(s))
}

// NextIntWithError 读取下一个整数，若出错则返回该错误。
func NextIntWithError() (int, error) {
	s := NextString()
	return strconv.Atoi(s)
}

// NextFloat64WithError 读取下一个 float64，若出错则返回该错误。
func NextFloat64WithError() (float64, error) {
	s := NextString()
	return strconv.ParseFloat(s, 64)
}

// NextIntSlice 读取 n 个整数并返回整数切片。
func NextIntSlice(n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = NextInt()
	}
	return res
}

// NewIncreaseIntSlice 创建长度为 n 的递增整数切片，首元素为 base，即 [base, base+1, ..., base+n-1]。
func NewIncreaseIntSlice(n int, base int) []int {
	newSlice := make([]int, n)
	for i := 0; i < n; i++ {
		newSlice[i] = i + base
	}
	return newSlice
}

// Next2IntSlice 读取 n 对整数，分别返回两个长度为 n 的切片 a 和 b。
func Next2IntSlice(n int) ([]int, []int) {
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = NextInt2()
	}
	return a, b
}

// NextIntSlice2D 读取一个 n×m 的二维整数矩阵。
func NextIntSlice2D(n, m int) [][]int {
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = NextIntSlice(m)
	}
	return a
}

// NewIntSlice2D 创建一个 n×m 的二维整数切片，所有元素初始化为 def。
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

// NextFloat 读取下一个 float64。若无法解析则 panic。
func NextFloat() float64 {
	return unwrap(strconv.ParseFloat(NextString(), 64))
}

// NextFloats 读取 n 个 float64 并返回切片。
func NextFloats(n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = NextFloat()
	}
	return res
}

// NextInt2 连续读取两个整数并返回。
func NextInt2() (int, int) {
	return NextInt(), NextInt()
}

// NextInt3 连续读取三个整数并返回。
func NextInt3() (int, int, int) {
	return NextInt(), NextInt(), NextInt()
}

// NextInt4 连续读取四个整数并返回。
func NextInt4() (int, int, int, int) {
	return NextInt(), NextInt(), NextInt(), NextInt()
}

// Print 将 a 写入缓冲输出，不自动添加换行。
func Print(a ...any) {
	if _, err := fmt.Fprint(wr, a...); err != nil {
		panic(fmt.Errorf("print: %w", err))
	}
}

// Printf 按指定格式写入缓冲输出。
func Printf(format string, a ...any) {
	fmt.Fprintf(wr, format, a...)
}

// Println 将 a 写入缓冲输出并追加换行。
func Println(a ...any) {
	fmt.Fprintln(wr, a...)
}

// PrintSlice 以空格分隔打印切片中的每个元素，不追加换行。
func PrintSlice[T constraints.Ordered](slices []T) {
	for _, v := range slices {
		Print(v, " ")
	}
}

// PrintlnSlice 逐行打印切片中的每个元素。
func PrintlnSlice[T constraints.Ordered](slices []T) {
	for _, v := range slices {
		Println(v)
	}
}

// PrintlnSliceInline 以空格分隔打印切片，末尾输出换行。
func PrintlnSliceInline[T constraints.Ordered](slices []T) {
	if len(slices) > 0 {
		Print(fmt.Sprintf("%v", slices[0]))
		for _, v := range slices[1:] {
			Print(fmt.Sprintf(" %v", v))
		}
	}
	Println()
}

// Flush 刷新缓冲写入器，将缓冲内容真正输出。
func Flush() {
	wr.Flush()
}

// Debug 仅在 debug 模式（os.Args[1] == "i"）下输出内容，否则不做任何事。
func Debug(v ...interface{}) {
	if !debugFlg {
		return
	}
	Println(v...)
}
