// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 4.8: Modify charcount to count letters, digits, and so on in their
// Unicode categories, using functions like unicode.IsLetter.

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

var categories = []struct {
	name string
	fn   func(rune) bool
}{
	{"control", unicode.IsControl},
	{"digit", unicode.IsDigit},
	{"graphic", unicode.IsGraphic},
	{"letter", unicode.IsLetter},
	{"lower", unicode.IsLower},
	{"mark", unicode.IsMark},
	{"number", unicode.IsNumber},
	{"print", unicode.IsPrint},
	{"punct", unicode.IsPunct},
	{"space", unicode.IsSpace},
	{"symbol", unicode.IsSymbol},
	{"title", unicode.IsTitle},
	{"upper", unicode.IsUpper},
}

func categorize(r rune, m map[string]int) {
	for _, cat := range categories {
		if cat.fn(r) {
			m[cat.name]++
		}
	}
}

func main() {
	runecounts := make(map[rune]int)  // counts of Unicode characters
	catcounts := make(map[string]int) // counts of Unicode categories
	var utflen [utf8.UTFMax + 1]int   // count of lengths of UTF-8 encodings
	invalid := 0                      // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		runecounts[r]++
		utflen[n]++
		categorize(r, catcounts)
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range runecounts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\ncategory\tcount\n")
	for c, n := range catcounts {
		fmt.Printf("%-15s\t%d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//!-
