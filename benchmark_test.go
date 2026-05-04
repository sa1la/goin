package goin

import (
	"testing"
)

func BenchmarkPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pow(2, 10)
	}
}

func BenchmarkModPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ModPow(2, 100, 1000000007)
	}
}

func BenchmarkGcd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Gcd(1234567, 987654)
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(982451653)
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(30)
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(10)
	}
}

func BenchmarkStackPush(b *testing.B) {
	stack := &Stack[int]{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkStackPop(b *testing.B) {
	stack := &Stack[int]{}
	// Pre-fill stack
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Pop()
	}
}

func BenchmarkQueueEnqueue(b *testing.B) {
	queue := NewQueue[int](1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
	}
}

func BenchmarkQueueDequeue(b *testing.B) {
	queue := NewQueue[int](b.N)
	// Pre-fill queue
	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queue.Dequeue()
	}
}

func BenchmarkMinHeapPush(b *testing.B) {
	heap := NewMinHeap[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.Push(i)
	}
}

func BenchmarkMinHeapPop(b *testing.B) {
	heap := NewMinHeap[int]()
	// Pre-fill heap
	for i := 0; i < b.N; i++ {
		heap.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.Pop()
	}
}

func BenchmarkAVLTreeInsert(b *testing.B) {
	tree := &AVLTree{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Insert(i)
	}
}

func BenchmarkAVLTreeSearch(b *testing.B) {
	tree := &AVLTree{}
	// Pre-fill tree
	for i := 0; i < 1000; i++ {
		tree.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Search(i % 1000)
	}
}

func BenchmarkNextPermutation(b *testing.B) {
	data := []int{1, 2, 3, 4, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NextPermutation(data)
	}
}

func BenchmarkSieveOfEratosthenes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SieveOfEratosthenes(1000)
	}
}

func BenchmarkCombo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Combo(20, 10)
	}
}

func BenchmarkUnionFindUnion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		uf := NewUnionFind(1000)
		b.StartTimer()
		uf.Union(i%1000, (i*7+13)%1000)
	}
}

func BenchmarkUnionFindFind(b *testing.B) {
	uf := NewUnionFind(1000)
	for i := 0; i < 999; i++ {
		uf.Union(i, i+1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uf.Find(i % 1000)
	}
}

func BenchmarkFenwickTreeUpdate(b *testing.B) {
	ft := NewFenwickTree[int](10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ft.Update(i%10000, 1)
	}
}

func BenchmarkFenwickTreeQuery(b *testing.B) {
	ft := NewFenwickTree[int](10000)
	for i := 0; i < 10000; i++ {
		ft.Update(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ft.Query(i % 10000)
	}
}

func BenchmarkFenwickTreeFromSlice(b *testing.B) {
	a := make([]int, 100000)
	for i := range a {
		a[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewFenwickTreeFromSlice(a)
	}
}
