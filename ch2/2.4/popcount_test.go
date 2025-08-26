// Benchmarks for exercise 2.4

/*
dibble@fedora:~/learn/gobook/ch2/2.4$ go test -bench=.
goos: linux
goarch: amd64
pkg: gobook/ch2/2.4
cpu: AMD Ryzen 7 7840U w/ Radeon  780M Graphics
BenchmarkPopcount-16         	658963792	         1.806 ns/op
BenchmarkPopcountShift-16    	57607233	        18.12 ns/op
PASS
ok  	gobook/ch2/2.4	2.238s
*/

package popcount_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	pc1 "gobook/ch2/2.3"
	pc2 "gobook/ch2/2.4"
)

func BenchmarkPopcount(b *testing.B) {
	for b.Loop() {
		pc1.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopcountShift(b *testing.B) {
	for b.Loop() {
		pc2.PopCountShift(0x1234567890ABCDEF)
	}
}

func TestPopcount(t *testing.T) {
	for i := range uint64(999) {
		assert.Equal(t, pc1.PopCount(i), pc2.PopCountShift(i))
	}
}
