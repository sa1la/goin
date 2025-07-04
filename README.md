# Goin - Go 竞技编程工具库

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.22-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

一个专为竞技编程设计的 Go 语言工具库，提供高效、易用的数据结构和算法实现。

## ✨ 特性

- 🚀 **高性能**: 针对竞技编程优化的实现
- 🔧 **易用性**: 简洁直观的 API 设计
- 📦 **零依赖**: 除了测试库，无其他外部依赖
- 🧪 **高测试覆盖**: 完善的单元测试保证代码质量
- 🎯 **泛型支持**: 充分利用 Go 1.18+ 的泛型特性

## 🚀 快速开始

```bash
go get github.com/sa1la/goin
```

```go
package main

import "github.com/sa1la/goin"

func main() {
    // 读取输入
    n := goin.NextInt()
    arr := goin.NextIntSlice(n)
    
    // 数学计算
    result := goin.Pow(2, 10)
    gcd := goin.Gcd(12, 18)
    
    // 数据结构操作
    stack := &goin.Stack[int]{}
    stack.Push(42)
    
    // 输出结果
    goin.Println(result)
    goin.Flush()
}
```

## 📚 功能模块

### 🔢 输入输出 (IO)
- `NextInt()`, `NextString()` - 快速输入读取
- `NextIntSlice(n)` - 批量读取整数数组
- `Print()`, `Println()` - 高效输出
- `Debug()` - 调试输出支持

### 🧮 数学工具 (Math)
- `Pow(x, n)` - 快速幂算法 O(log n)
- `Gcd(a, b)` - 最大公约数
- `Lcm(a, b)` - 最小公倍数
- `IsPrime(n)` - 素数判断
- `Factorial(n)` - 阶乘计算
- `ModPow(base, exp, mod)` - 模快速幂
- `Fibonacci(n)` - 斐波那契数列
- `SieveOfEratosthenes(n)` - 埃拉托斯特尼筛法

### 📊 数据结构
#### 栈 (Stack)
```go
stack := &goin.Stack[int]{}
stack.Push(1)
val := stack.Pop()
```

#### 队列 (Queue)
```go
queue := goin.NewQueue[int](10)
queue.Enqueue(1)
val := queue.Dequeue()
```

#### 堆 (Heap)
```go
minHeap := goin.NewMinHeap[int]()
maxHeap := goin.NewMaxHeap[int]()
pq := goin.NewPriorityQueue[int]()
```

#### 二叉树
- `SliceBinaryTree[T]` - 基于数组的二叉树
- `AVLTree` - 自平衡二叉搜索树

### 🔄 算法工具
- `NextPermutation()` - 下一个排列
- `LastPermutation()` - 上一个排列
- `Reverse()` - 数组反转
- `FillSlice()` - 数组填充

## 📋 示例

### 基本 I/O 操作
```go
// 读取多个整数
a, b := goin.NextInt2()
x, y, z := goin.NextInt3()

// 读取字符串数组
strings := goin.NextStrings(3)

// 输出数组
arr := []int{1, 2, 3, 4, 5}
goin.PrintlnSliceInline(arr) // 输出: 1 2 3 4 5
```

### 数学计算
```go
// 计算组合数 C(n, k)
result := goin.Combo(10, 3) // 120

// 生成素数
primes := goin.SieveOfEratosthenes(100)

// 模运算
modResult := goin.ModPow(2, 100, 1000000007)
```

### 数据结构使用
```go
// 使用优先队列
pq := goin.NewPriorityQueue[int]()
pq.Enqueue(3)
pq.Enqueue(1)
pq.Enqueue(4)
fmt.Println(pq.Dequeue()) // 输出: 4 (最大值)

// AVL 树操作
tree := &goin.AVLTree{}
tree.Insert(10)
tree.Insert(5)
tree.Insert(15)
inOrder := tree.InOrder() // [5, 10, 15]
```

## 🔧 调试支持

设置调试模式，从文件读取输入：
```bash
go run main.go i  # 从 ./input 文件读取输入
```

## 🧪 运行测试

```bash
go test -v ./...
```

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📞 支持

如果您在使用过程中遇到问题，请：
1. 查看现有的 [Issues](https://github.com/sa1la/goin/issues)
2. 创建新的 Issue 描述问题
3. 提供最小可复现示例

---

⭐ 如果这个项目对您有帮助，请给个 Star！
