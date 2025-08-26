// Benchmarks for exercise 2.5

/*
dibble@fedora:~/learn/gobook/ch2/2.5$ go test -bench=.
goos: linux
goarch: amd64
pkg: gobook/ch2/2.5
cpu: AMD Ryzen 7 7840U w/ Radeon  780M Graphics
BenchmarkPopcount-16           	655279527	         1.818 ns/op
BenchmarkPopcountSpecial-16    	140192066	         8.584 ns/op
PASS
ok  	gobook/ch2/2.5	2.399s
*/

package popcount_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	pc1 "gobook/ch2/2.3"
	pc3 "gobook/ch2/2.5"
)

func BenchmarkPopcount(b *testing.B) {
	for b.Loop() {
		pc1.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopcountSpecial(b *testing.B) {
	for b.Loop() {
		pc3.PopCountSpecial(0x1234567890ABCDEF)
	}
}

func TestPopcount(t *testing.T) {
	for i := range uint64(999) {
		assert.Equal(t, pc1.PopCount(i), pc3.PopCountSpecial(i))
	}
}
