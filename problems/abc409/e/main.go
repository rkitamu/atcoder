package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	// "container/heap"
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

const FACTORIAL_CACHE_SIZE = 10000000
const MOD = 1000000007

func main() {
	defer flush()
	n := ni()
	g := NewGraph()

	for i := 1; i <= n; i++ {
		val := ni()
		g.AddNode(i, val)
	}

	for i := 0; i < n-1; i++ {
		u := ni()
		v := ni()
		w := ni()
		g.AddUndirectedEdge(u, v, w)
	}

	var ans int64
	var dfs func(curr, parent int) int
	dfs = func(curr, parent int) int {
		node := g.GetNode(curr)
		total := node.Value

		for _, edge := range node.Edges {
			if edge.To == parent {
				continue
			}
			childTotal := dfs(edge.To, curr)
			ans += int64(abs(childTotal)) * int64(edge.Weight)
			total += childTotal
		}

		return total
	}

	dfs(1, -1)
	out(ans)
}

// =====================
// utils
// =====================
// io
var sc = bufio.NewScanner(os.Stdin)
var rdr = bufio.NewReader(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

// =====================
// input utils
// =====================
// ni reads a single integer from stdin.
func ni() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

// ni reads n integers from stdin.
func nis(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ni()
	}
	return a
}

var bufBytes []byte
var bufIdx int

func nr() byte {
	for {
		if bufIdx < len(bufBytes) {
			r := bufBytes[bufIdx]
			bufIdx++
			return r
		}

		if !sc.Scan() {
			panic("failed to scan next token")
		}
		bufBytes = []byte(sc.Text())
		bufIdx = 0
	}
}

// nr reads a single string from stdin.
func ns() string {
	if !sc.Scan() {
		panic("failed to scan next token")
	}
	return sc.Text()
}

/* なんかtest.sh実行時だけエラーでる
// nr reads a single rune from stdin.
func nr() rune {
	for {
		r, _, err := rdr.ReadRune()
		if err != nil {
			panic(err)
		}
		if r != '\n' && r != '\r' {
			return r
		}
	}
}*/

// nr reads n bytes from stdin.
func nrs(n int) []byte {
	a := make([]byte, n)
	for i := 0; i < n; i++ {
		a[i] = nr()
	}
	return a
}

// nrs2d reads n * m bytes from stdin.
func nrs2d(n, m, offset int) [][]byte {
	a := make([][]byte, n+offset)
	for i := offset; i < n+offset; i++ {
		tmp := nrs(m)
		prepended := make([]byte, offset+len(tmp))
		copy(prepended[offset:], tmp)
		a[i] = prepended
	}
	return a
}

// =====================
// output utils
// =====================
// flush flushes the buffered writer.
func flush() {
	e := wtr.Flush()
	if e != nil {
		panic(e)
	}
}

// 相対誤差が10^-6: formatFloat(f, 7)
func formatFloat(f float64, precision int) string {
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return fmt.Sprintf("%v", f)
	}

	magnitude := math.Log10(math.Abs(f))
	effectivePrecision := precision - int(magnitude) - 1

	if effectivePrecision < 0 {
		effectivePrecision = 0
	}
	if effectivePrecision > 15 { // float64の限界
		effectivePrecision = 15
	}

	return fmt.Sprintf("%.*f", effectivePrecision, f)
}

// out writes the output to stdout.
func out(v ...interface{}) {
	_, e := fmt.Fprintln(wtr, v...)
	if e != nil {
		panic(e)
	}
}

// outr2d writes a 2D slice of runes to stdout.
func outr2d(a [][]rune) {
	for _, r := range a {
		_, e := fmt.Fprintln(wtr, string(r))
		if e != nil {
			panic(e)
		}
	}
}

// ======================
// type utils
// ======================
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// =====================
// calculation utils
// =====================
var fact, invFact []int
var factorialInitialized = false

// initFactorialTable initializes the factorial cache table
func initFactorialTable() {
	if factorialInitialized {
		return
	}
	factorialInitialized = true

	fact = make([]int, FACTORIAL_CACHE_SIZE+1)
	invFact = make([]int, FACTORIAL_CACHE_SIZE+1)
	fact[0] = 1
	for i := 1; i <= FACTORIAL_CACHE_SIZE; i++ {
		fact[i] = fact[i-1] * i % MOD
	}
	invFact[FACTORIAL_CACHE_SIZE] = powMod(fact[FACTORIAL_CACHE_SIZE], MOD-2)
	for i := FACTORIAL_CACHE_SIZE - 1; i >= 0; i-- {
		invFact[i] = invFact[i+1] * (i + 1) % MOD
	}
}

// combination calculates nCk
func combMod(n, k int) int {
	initFactorialTable()
	if n < 0 || k < 0 || k > n {
		return 0
	}
	return (fact[n] * invFact[k] % MOD * invFact[n-k] % MOD) % MOD
}

func powMod(x, e int) int {
	res := 1
	for e > 0 {
		if e%2 == 1 {
			res = res * x % MOD
		}
		x = x * x % MOD
		e /= 2
	}
	return res
}

// ======================
// data structure
// ======================
// BIT is a Binary Indexed Tree (Fenwick Tree) implementation
type BIT struct {
	n   int
	bit []int
}

func NewBIT(n int) *BIT {
	return &BIT{n: n + 2, bit: make([]int, n+3)}
}

func (b *BIT) Add(i, x int) {
	i++
	for i < len(b.bit) {
		b.bit[i] += x
		i += i & -i
	}
}

func (b *BIT) Sum(i int) int {
	i++
	res := 0
	for i > 0 {
		res += b.bit[i]
		i -= i & -i
	}
	return res
}

// SegmentTree (WIP(Implemented: push, add, get, sum))
type SegmentTree struct {
	n    int
	data []int
	lazy []int
}

func NewSegmentTree(n int) *SegmentTree {
	size := 1
	for size < n {
		size <<= 1
	}
	return &SegmentTree{
		n:    size,
		data: make([]int, 2*size),
		lazy: make([]int, 2*size),
	}
}

func (st *SegmentTree) push(k, l, r int) {
	if st.lazy[k] != 0 {
		st.data[k] += (r - l) * st.lazy[k]
		if r-l > 1 {
			st.lazy[2*k+1] += st.lazy[k]
			st.lazy[2*k+2] += st.lazy[k]
		}
		st.lazy[k] = 0
	}
}

// Add x to range [a, b)
func (st *SegmentTree) Add(a, b, x int) {
	var f func(k, l, r int)
	f = func(k, l, r int) {
		st.push(k, l, r)
		if r <= a || b <= l {
			return
		}
		if a <= l && r <= b {
			st.lazy[k] += x
			st.push(k, l, r)
		} else {
			mid := (l + r) / 2
			f(2*k+1, l, mid)
			f(2*k+2, mid, r)
			st.data[k] = st.data[2*k+1] + st.data[2*k+2]
		}
	}
	f(0, 0, st.n)
}

// Get value at index i
func (st *SegmentTree) Get(i int) int {
	k := 0
	l, r := 0, st.n
	for r-l > 1 {
		st.push(k, l, r)
		mid := (l + r) / 2
		if i < mid {
			k = 2*k + 1
			r = mid
		} else {
			k = 2*k + 2
			l = mid
		}
	}
	st.push(k, l, r)
	return st.data[k]
}

// Get sum of range [a, b)
func (st *SegmentTree) Sum(a, b int) int {
	var f func(k, l, r int) int
	f = func(k, l, r int) int {
		st.push(k, l, r)
		if r <= a || b <= l {
			return 0
		}
		if a <= l && r <= b {
			return st.data[k]
		}
		mid := (l + r) / 2
		return f(2*k+1, l, mid) + f(2*k+2, mid, r)
	}
	return f(0, 0, st.n)
}

// Stack is a simple stack implementation
type Stack[T any] struct {
	data []T
}

func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{data: make([]T, size)}
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() T {
	last := len(s.data) - 1
	v := s.data[last]
	s.data = s.data[:last]
	return v
}

func (s *Stack[T]) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Len() int {
	return len(s.data)
}

func (s *Stack[T]) Top() T {
	return s.data[len(s.data)-1]
}

// Queue is a simple queue implementation
type Queue[T any] struct {
	data []T
	head int
	tail int
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{data: make([]T, size), head: 0, tail: 0}
}
func (q *Queue[T]) Enqueue(v T) {
	q.data = append(q.data, v)
	q.tail++
}
func (q *Queue[T]) Dequeue() T {
	if q.head == q.tail {
		panic("queue is empty")
	}
	v := q.data[q.head]
	q.head++
	if q.head == len(q.data)/2 {
		q.data = q.data[q.head:]
		q.tail -= q.head
		q.head = 0
	}
	return v
}
func (q *Queue[T]) Empty() bool {
	return q.head == q.tail
}
func (q *Queue[T]) Len() int {
	return q.tail - q.head
}
func (q *Queue[T]) Top() T {
	if q.head == q.tail {
		panic("queue is empty")
	}
	return q.data[q.head]
}

// Priority Queue
// usage:
//
//	import "container/heap"
//	h := &ItemHeap{}
//	heap.Init(h)
//	heap.Push(h, &Item{value: tc.tcase[i]})
//	heap.Pop(h).(*Item)
type Item struct {
	value int
}
type ItemHeap []*Item

func (h ItemHeap) Len() int { return len(h) }

// min-heap implementation
func (h ItemHeap) Less(i, j int) bool  { return h[i].value < h[j].value }
func (h ItemHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *ItemHeap) Push(x interface{}) { *h = append(*h, x.(*Item)) }
func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Vector struct {
	X, Y int
}

func NewVector(x, y int) *Vector {
	return &Vector{X: x, Y: y}
}

func NewVectorFromPointsSlice(start, end []int) *Vector {
	if len(start) < 2 || len(end) < 2 {
		panic("require at least 2 elements (x, y)")
	}
	return &Vector{
		X: end[0] - start[0],
		Y: end[1] - start[1],
	}
}

func (v Vector) Add(other Vector) *Vector {
	return &Vector{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func (v Vector) Dot(other Vector) float64 {
	return float64(v.X*other.X + v.Y*other.Y)
}

func (v Vector) Cross(other Vector) int {
	return v.X*other.Y - v.Y*other.X
}

func (v Vector) CrossMagnitude(other Vector) float64 {
	return math.Abs(float64(v.Cross(other)))
}

// Edge represents a connection from one node to another with an optional weight.
type Edge struct {
	To     int
	Weight int
}

// Node represents a node in the graph with optional value and its outgoing edges.
type Node struct {
	ID    int
	Value int
	Edges []Edge
}

// Graph represents a generic directed or undirected graph.
type Graph struct {
	Nodes map[int]*Node
}

// NewGraph initializes an empty graph.
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[int]*Node),
	}
}

// AddNode adds a new node with the given ID and optional value.
func (g *Graph) AddNode(id int, value int) {
	g.Nodes[id] = &Node{
		ID:    id,
		Value: value,
		Edges: []Edge{},
	}
}

// AddEdge adds a directed edge from u to v with given weight.
func (g *Graph) AddEdge(u, v, weight int) {
	if _, ok := g.Nodes[u]; !ok {
		g.AddNode(u, 0)
	}
	if _, ok := g.Nodes[v]; !ok {
		g.AddNode(v, 0)
	}
	g.Nodes[u].Edges = append(g.Nodes[u].Edges, Edge{To: v, Weight: weight})
}

func (g *Graph) AddUndirectedEdge(u, v, weight int) {
	g.AddEdge(u, v, weight)
	g.AddEdge(v, u, weight)
}

// GetNode returns the node with the given ID.
func (g *Graph) GetNode(id int) *Node {
	return g.Nodes[id]
}

// =====================
// Math utils
// ======================
// isPrime checks if n is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	cur := 3
	max := int(math.Floor(float64(math.Sqrt(float64(n)))))
	for cur <= max {
		if m := n % cur; m == 0 {
			return false
		}
		cur++
	}
	return true
}

// gcd calculates the greatest common divisor of a and b
func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for 1 <= a && 1 <= b {
		mod := a % b
		if mod == 0 {
			return b
		}
		a, b = b, mod
	}
	if 1 <= a {
		return a
	}
	return b
}

// gcds calculates the greatest common divisor of a slice of integers
func gcds(a ...int) int {
	if len(a) < 2 {
		panic("gcds: at least 2 arguments required")
	}
	g := a[0]
	for i := 1; i < len(a); i++ {
		g = gcd(g, a[i])
	}
	return g
}

// lcm calculates the least common multiple of a and b
func lcm(a, b int) int {
	if a < b {
		a, b = b, a
	}
	return a / gcd(a, b) * b
}

func lcms(a ...int) int {
	if len(a) < 2 {
		panic("lcms: at least 2 arguments required")
	}
	l := a[0]
	for i := 1; i < len(a); i++ {
		l = lcm(l, a[i])
	}
	return l
}

var factorialCache = make([]int64, 0)

func factorial(n int) int64 {
	if n < 0 {
		panic("factorial: n must be non-negative")
	}
	if len(factorialCache) > n {
		return factorialCache[n]
	}
	for i := len(factorialCache); i <= n; i++ {
		if i == 0 {
			factorialCache = append(factorialCache, 1)
		} else {
			factorialCache = append(factorialCache, factorialCache[i-1]*int64(i))
		}
	}
	return factorialCache[n]
}

func comb(n, k int) int64 {
	if n < 0 || k < 0 || k > n {
		return 0
	}
	return factorial(n) / (factorial(k) * factorial(n-k))
}

func min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func abs[T Number](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

func fibonacci(n int) int {
	if n < 0 {
		panic("fibonacci: n must be non-negative")
	}
	if n == 0 {
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

// =====================
// Programming utils
// ======================
// Pair is C++のPairっぽいやつ
// WIP: Number以外(constraints.Ordered にする)
type Pair[T, U Number] struct {
	First  T
	Second U
}

func NewPair[T, U Number](first T, second U) Pair[T, U] {
	return Pair[T, U]{First: first, Second: second}
}

func (p Pair[T, U]) Lt(other Pair[T, U]) bool {
	if p.First == other.First {
		return p.Second < other.Second
	} else {
		return p.First < other.First
	}
}

func (p Pair[T, U]) Lte(other Pair[T, U]) bool {
	if p.First == other.First {
		return p.Second <= other.Second
	} else {
		return p.First <= other.First
	}
}

func (p Pair[T, U]) Gt(other Pair[T, U]) bool {
	if p.First == other.First {
		return p.Second > other.Second
	} else {
		return p.First > other.First
	}
}

func (p Pair[T, U]) Gte(other Pair[T, U]) bool {
	if p.First == other.First {
		return p.Second >= other.Second
	} else {
		return p.First >= other.First
	}
}

func (p Pair[T, U]) Max(other Pair[T, U]) Pair[T, U] {
	if p.Lt(other) {
		return other
	}
	return p
}

func (p Pair[T, U]) Min(other Pair[T, U]) Pair[T, U] {
	if p.Lt(other) {
		return p
	}
	return other
}

// Swap swaps the values of two variables.
func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}
