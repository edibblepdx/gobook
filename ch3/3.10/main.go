// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.10: Make comma iterative and use bytes.Buffer.

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// !+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	// You should only use decimals, but if some crazy
	// person put runes into numbers then it'll still
	// work. This is stupid though :)

	sr := []rune(s)
	slen := len(sr) - 1
	var buf bytes.Buffer

	for i, r := range sr {
		buf.WriteRune(r)
		if slen != i && (slen-i)%3 == 0 {
			buf.WriteString(",")
		}
	}

	return buf.String()
}

//!-
