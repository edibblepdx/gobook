// Exercise 4.1: Write a program to count the number of bits that are different
// in two SHA256 hashes.

package main

import (
	"crypto/sha256"
	"fmt"
)

func sha256BitDiff(h1, h2 [32]byte) int {
	var count int
	for i := range 32 {
		b1, b2 := h1[i], h2[i]
		for j := range 8 {
			if (b1>>j)&1 != (b2>>j)&1 {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(sha256BitDiff(sha256.Sum256([]byte("x")), sha256.Sum256([]byte("X"))))
}
