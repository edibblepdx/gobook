// Exercise 4.4: Write a version of rotate that operates in a single pass.
// Usage: go run . <length of array> <shift>

package main

import (
	"fmt"
	"os"
	"strconv"
)

// Greatest common divisor
// Panics if either m or n are zero
func gcd(m, n int) int {
	d := min(m, n)
	for d > 0 {
		if m%d == 0 && n%d == 0 {
			break
		}
		d--
	}
	return d
}

// Absolute value of an integer
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// The modulo function defined in terms of floor division
// https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/divmodnote-letter.pdf
func modF(D, d int) int {
	r := D % d
	if (r > 0 && d < 0) || (r < 0 && d > 0) {
		r += d
	}
	return r
}

// Rotate an array n spaces in place
func rotate(s []int, n int) {
	l := len(s)
	if l == 0 {
		return
	}

	n %= l
	if n == 0 {
		return
	}

	start := -1
	for range gcd(l, abs(n)) {
		start++
		idx, store := start, s[start]
		for {
			idx = modF(idx-n, l) // shift & wrap
			s[idx], store = store, s[idx]

			if idx == start {
				break
			}
		}
	}
}

func try(m, n int) {
	a := make([]int, m)
	for i := range m {
		a[i] = i + 1
	}

	rotate(a, n)
	fmt.Println(a, n)
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: go run . <length of array> <shift>")
		os.Exit(1)
	}

	m, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}

	n, err := strconv.Atoi(args[2])
	if err != nil {
		panic(err)
	}

	try(m, n)
}
