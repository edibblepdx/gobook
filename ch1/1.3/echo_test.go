// Exercise 1.3: Measure the difference in running time between potentially
// slow echo implementations and one that uses strings.Join

/*
$ go test . -bench=.
goos: linux
goarch: amd64
pkg: gobook/ch1/1.3
cpu: AMD Ryzen 7 7840U w/ Radeon  780M Graphics
BenchmarkLoopVariant-16           	19809876	        60.69 ns/op
BenchmarkStringsJoinVariant-16    	32248519	        37.46 ns/op
PASS
ok  	gobook/ch1/1.3	2.413s
*/

package main

import (
	"os"
	"strings"
	"testing"
)

func BenchmarkLoopVariant(b *testing.B) {
	for b.Loop() {
		var s, sep string
		for _, val := range os.Args[1:] {
			s += val + sep
			sep = " "
		}
	}
}

func BenchmarkStringsJoinVariant(b *testing.B) {
	for b.Loop() {
		strings.Join(os.Args[1:], " ")
	}
}
