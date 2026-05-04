# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

`github.com/sa1la/goin` — a Go competitive-programming toolbox. Single flat package (`package goin`) sitting at repo root; users import the whole thing into a contest template. Go 1.22+ (toolchain 1.23.2), generics-heavy via `golang.org/x/exp/constraints`.

## Common commands

```bash
# Run all tests (verbose)
go test -v ./...

# Run a single test by name (regex)
go test -run TestMinHeap -v

# Run all tests in one file
go test -run TestNextString -v        # io_test.go contains TestNextString*
# (Go test takes regexes, not filenames — match by test name prefix)

# Run benchmarks
go test -bench=. -benchmem
go test -bench=BenchmarkPow -benchmem  # single benchmark

# Race detector
go test -race ./...

# Coverage
go test -cover ./...
go test -coverprofile=cover.out && go tool cover -html=cover.out

# Build / vet / format
go build ./...
go vet ./...
gofmt -w .

# Debug input mode (reads ./input as stdin — see io.go init())
go run yourmain.go i
```

## Architecture

Flat package, one file per concern. There is no `main` — this is a library only.

| File | Responsibility |
|------|----------------|
| `io.go` | Fast scanner/writer, input parsers (`NextInt`, `NextIntSlice`, …), output (`Print`, `Println`, `PrintlnSliceInline`), `Debug`, `Flush` |
| `math.go` | `Max`/`Min`/`ChMax`/`ChMin`, `Abs`, `Pow`, `Gcd`/`Lcm`, `IsPrime`, `Factorial`, `ModPow`, `ModInv`, `ExtendedGcd`, `SieveOfEratosthenes`, `Fibonacci`, `Combo`, `IsPowerOfTwo`/`NextPowerOfTwo`, `GetAngle`, `Sum`, plus `const Inf = math.MaxInt64` |
| `slice.go` | `FillSlice`, `Reverse`, `NextPermutation`, `LastPermutation` |
| `stack.go` | `Stack[T any]` — generic LIFO over a slice |
| `queue.go` | `Queue[T any]` — circular buffer FIFO with `resize()` doubling |
| `heap.go` | `Heap[T constraints.Ordered]` with `MinHeap`/`MaxHeap` discriminator; `PriorityQueue[T]` is a thin `Enqueue`/`Dequeue` alias over a **max-heap** |
| `tree.go` | `SliceBinaryTree[T]` (heap-style array tree, level-order ↔ index math) and `AVLTree` (int-only self-balancing BST with rotations) |

### Cross-cutting things to know

- **Global I/O state.** `io.go` declares package-level `sc *bufio.Scanner` and `wr *bufio.Writer`. They're configured in `init()`: `sc` uses `ScanWords` with an unbounded buffer, `wr` writes to stdout. Anything calling `Print*` MUST end with `Flush()` or output is lost. Tests reassign `sc` to a `strings.NewReader` to inject input — keep that pattern when adding I/O tests.
- **Debug mode.** If `os.Args[1] == "i"`, `init()` reads `./input` and replaces spaces with newlines before feeding the scanner (because the scanner splits on words anyway, this just lets humans write space-separated test data). Sets `debugFlg = true`, which gates `Debug(...)` output.
- **Generics constraint.** Most container/slice operations use `constraints.Ordered` (from `golang.org/x/exp/constraints`) rather than `cmp.Ordered`. Don't switch to `cmp.Ordered` without bumping the Go floor and checking call sites.
- **`AVLTree` is int-only.** Unlike `Heap`/`Stack`/`Queue`, `AVLTree` operates on `int` directly (`TreeNode.Val int`). Genericizing it would touch rotations, comparisons, and the DFS helpers — non-trivial.
- **`PriorityQueue` defaults to max-heap.** `NewPriorityQueue[T]()` wraps `NewMaxHeap[T]()`. `Dequeue` returns the largest element. If you need a min-priority-queue, use `NewMinHeap` directly.
- **Zero-value pop.** `Stack.Pop`, `Queue.Dequeue`, `Heap.Pop` all return `var zero T` on empty rather than panicking. Callers needing to distinguish "empty" from "popped zero" must `IsEmpty()` / check `Size()` first; `Peek` returns `(T, bool)` for that reason.

## Test conventions

- Tests live next to sources (`<name>_test.go`), all in `package goin`.
- Use `github.com/stretchr/testify/assert` for assertions; the older `testing.T.Errorf` style appears only in `io_test.go` because it predates testify adoption — match testify in new tests.
- Benchmarks are consolidated in `benchmark_test.go`. When adding a feature with a hot path, add a `BenchmarkX` there rather than scattering benchmarks across files.
- I/O tests **must** rebind `sc` (e.g. `sc = bufio.NewScanner(strings.NewReader(...))`) before calling `Next*`. Don't rely on whatever the previous test left in the global.
