package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	// if use Combination
	// initCombTable()
}

const MAX_COMB_CACHE = 10000000
const MOD = 1000000007



func main() {
	defer flush()

	nm := nis(2)
	n, m := nm[0], nm[1]

	type Segment struct {
		l, r int
	}
	seg := make([]Segment, m)
	for i := 0; i < m; i++ {
		lr := nis(2)
		seg[i] = Segment{l: lr[0], r: lr[1]}
	}
	sort.Slice(seg, func(i, j int) bool { return seg[i].l < seg[j].l })

	q := ni()
	query := make([]*Segment, q)
	for i := range query {
		lr := nis(2)
		query[i] = &Segment{l: lr[0], r: lr[1]}
	}
	sort.Slice(query, func(i, j int) bool { return query[i].l < query[j].l })

	bit := NewBIT(2*n)
	res := make([]int, q)

	si := 0
	for _, q := range query {
		for si < m && seg[si].l <= q.l {
			bit.Add(seg[si].l, 1)
			si++
		}
		res[1] = bit.Sum(q.r) - bit.Sum(q.l-1)
	}
	
}

// =====================
// Portions of this file are based on code from:
// https://github.com/gosagawa/atcoder
// Copyright (c) gosagawa
// Licensed under the MIT License
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

var bufRunes []rune
var bufIdx int

func nr() rune {
	for {
		if bufIdx < len(bufRunes) {
			r := bufRunes[bufIdx]
			bufIdx++
			return r
		}

		if !sc.Scan() {
			panic("failed to scan next token")
		}
		bufRunes = []rune(sc.Text())
		bufIdx = 0
	}
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

// nr reads n runes from stdin.
func nrs(n int) []rune {
	a := make([]rune, n)
	for i := 0; i < n; i++ {
		a[i] = nr()
	}
	return a
}

// nrs2d reads n * m runes from stdin.
func nrs2d(n, m int) [][]rune {
	a := make([][]rune, n)
	for i := 0; i < n; i++ {
		a[i] = nrs(m)
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

// =====================
// calculation utils
// =====================
var fact, invFact []int

// combination calculates nCk
func initCombTable() {
	fact = make([]int, MAX_COMB_CACHE+1)
	invFact = make([]int, MAX_COMB_CACHE+1)
	fact[0] = 1
	for i := 1; i <= MAX_COMB_CACHE; i++ {
		fact[i] = fact[i-1] * i % MOD
	}
	invFact[MAX_COMB_CACHE] = powMod(fact[MAX_COMB_CACHE], MOD-2)
	for i := MAX_COMB_CACHE - 1; i >= 0; i-- {
		invFact[i] = invFact[i+1] * (i + 1) % MOD
	}
}

// comb calculates nCk
func combMod(n, k int) int {
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
type BIT struct {
	n int
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