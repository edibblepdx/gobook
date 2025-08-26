// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 2.5: Rewrite PopCount using the fact that x&(x-1) clears the
// rightmost non-zero bit

// See page 45.

// !+
package popcount

// PopCountSpecial returns the population count (number of set bits) of x.
func PopCountSpecial(x uint64) int {
	var popc int
	for ; x != 0; x &= x - 1 {
		popc++
	}
	return popc
}

//!-
