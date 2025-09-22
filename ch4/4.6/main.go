// Exercise 4.6: Write an in-place function that squashes each run of adjacent
// Unicode spaces in a UTF-8-encoded []byte slice into a single ASCII space.

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squash(b []byte) []byte {
	i, j := 0, 0
	sflag := false
	for i < len(b) {
		r, size := utf8.DecodeRune(b[i:])
		if unicode.IsSpace(r) {
			if !sflag {
				sflag = true
				b[j] = 32 // ascii space
				i, j = i+size, j+1
			}
			i += size
		} else {
			sflag = false
			copy(b[j:], b[i:i+size])
			i, j = i+size, j+size
		}
	}

	return b[:j]
}

func main() {
	b1 := []byte("\t\n\v\f\r\u0085\u00A0hi")
	fmt.Println(b1)
	b1 = squash(b1)
	fmt.Printf("%d\n%s\n\n", b1, b1)

	b2 := []byte("ab\t\n\vcd\f\ref\u0085\u00A0gh")
	fmt.Println(b2)
	b2 = squash(b2)
	fmt.Printf("%d\n%s\n\n", b2, b2)
}
