package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	defer flush()
	n := ni() // input single int
	_ = nis(n) // input n ints 

	out(n) // output result
}

// =====================
// utils
// =====================
// io
var sc = bufio.NewScanner(os.Stdin)
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
