// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.11: Write comma to deal with floating point numbers
// and optional sign.

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.

package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// !+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	// If there's no sign the index returned is -1
	// s[0:0] is just an empty string so this works.
	signIdx := strings.IndexAny(s, "+-")
	sign, s := s[:signIdx+1], s[signIdx+1:]

	intPart, fractPart := s, ""
	if dotIdx := strings.IndexByte(s, '.'); dotIdx >= 0 {
		intPart, fractPart = s[:dotIdx], s[dotIdx:]
	}

	var buf bytes.Buffer
	buf.WriteString(sign)

	intlen := len([]rune(intPart)) - 1
	for i, r := range intPart {
		buf.WriteRune(r)
		if intlen != i && (intlen-i)%3 == 0 {
			buf.WriteString(",")
		}
	}

	buf.WriteString(fractPart)

	return buf.String()
}

//!-
