/*
$ go test -bench . -benchmem
BenchmarkCopyWithLoop/10**1-4         	30000000	        39.9 ns/op	      80 B/op	       1 allocs/op
BenchmarkCopyWithLoop/10**2-4         	10000000	       187 ns/op	     896 B/op	       1 allocs/op
BenchmarkCopyWithLoop/10**3-4         	 1000000	      1479 ns/op	    8192 B/op	       1 allocs/op
BenchmarkCopyWithLoop/10**4-4         	  200000	     10776 ns/op	   81920 B/op	       1 allocs/op
BenchmarkCopyWithLoop/10**5-4         	   10000	    107451 ns/op	  802816 B/op	       1 allocs/op
BenchmarkCopyWithLoop/10**6-4         	    2000	   1069316 ns/op	 8003584 B/op	       1 allocs/op
BenchmarkCopyWithLoop/10**7-4         	     100	  11813525 ns/op	80003072 B/op	       1 allocs/op
BenchmarkCopyWithLoop/10**8-4         	      10	 136107449 ns/op	800006144 B/op	       1 allocs/op
BenchmarkCopyWithCopy/10**1-4         	30000000	        37.8 ns/op	      80 B/op	       1 allocs/op
BenchmarkCopyWithCopy/10**2-4         	10000000	       143 ns/op	     896 B/op	       1 allocs/op
BenchmarkCopyWithCopy/10**3-4         	 1000000	      1253 ns/op	    8192 B/op	       1 allocs/op
BenchmarkCopyWithCopy/10**4-4         	  200000	      9222 ns/op	   81920 B/op	       1 allocs/op
BenchmarkCopyWithCopy/10**5-4         	   20000	     99408 ns/op	  802816 B/op	       1 allocs/op
BenchmarkCopyWithCopy/10**6-4         	    2000	    972346 ns/op	 8003584 B/op	       1 allocs/op
BenchmarkCopyWithCopy/10**7-4         	     200	  10134779 ns/op	80003072 B/op	       1 allocs/op
BenchmarkCopyWithCopy/10**8-4         	      20	  99656721 ns/op	800006144 B/op	       1 allocs/op
*/
package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func CopyWithLoop(dst, src []float64) {
	for i, v := range src {
		dst[i] = v
	}
}

func CopyWithCopy(dst, src []float64) {
	copy(dst, src)
}

func BenchmarkCopyWithLoop(b *testing.B) {
	benchmarkCopy(b, CopyWithLoop)
}

func BenchmarkCopyWithCopy(b *testing.B) {
	benchmarkCopy(b, CopyWithCopy)
}

func benchmarkCopy(b *testing.B, f func([]float64, []float64)) {
	n := 1
	for x := 1; x <= 8; x++ {
		n *= 10
		src := makeSlice(n)
		b.Run(fmt.Sprintf("10**%d", x), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				dst := make([]float64, n)
				f(dst, src)
				if src[0] != dst[0] {
					b.Errorf("src[0] => %f, dst[0] => %f", src[0], dst[0])
				}
			}
		})
	}
}

func makeSlice(n int) []float64 {
	a := make([]float64, n)

	for i := 0; i < n; i++ {
		a[i] = rand.Float64()
	}

	return a
}
