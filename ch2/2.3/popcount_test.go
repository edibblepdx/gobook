// Benchmarks for exercise 2.3

/*
dibble@fedora:~/learn/gobook/ch2/2.3$ go test -bench=.
goos: linux
goarch: amd64
pkg: gobook/ch2/2.3
cpu: AMD Ryzen 7 7840U w/ Radeon  780M Graphics
BenchmarkPopcount-16        	654065398	         1.833 ns/op
BenchmarkPopcountLoop-16    	267914965	         4.448 ns/op
PASS
ok  	gobook/ch2/2.3	2.395s
*/

package popcount_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	pc "gobook/ch2/2.3"
)

func BenchmarkPopcount(b *testing.B) {
	for b.Loop() {
		pc.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopcountLoop(b *testing.B) {
	for b.Loop() {
		pc.PopCountLoop(0x1234567890ABCDEF)
	}
}

func TestPopcount(t *testing.T) {
	for i := range uint64(999) {
		assert.Equal(t, pc.PopCount(i), pc.PopCountLoop(i))
	}
}
