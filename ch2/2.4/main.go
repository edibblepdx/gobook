// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 2.4: Write PopCount with 64 shifts

// See page 45.

// !+
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountShift returns the population count (number of set bits) of x.
func PopCountShift(x uint64) int {
	var popc int
	for i := range 64 {
		popc += int((x >> i) & 1)
	}
	return popc
}

//!-
