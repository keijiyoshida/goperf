/*
$ go test -bench .
BenchmarkCopyWithLoop/10**1-4 	30000000	        38.3 ns/op
BenchmarkCopyWithLoop/10**2-4 	10000000	       184 ns/op
BenchmarkCopyWithLoop/10**3-4 	 1000000	      1528 ns/op
BenchmarkCopyWithLoop/10**4-4 	  100000	     10745 ns/op
BenchmarkCopyWithLoop/10**5-4 	   10000	    127916 ns/op
BenchmarkCopyWithLoop/10**6-4 	    2000	   1188813 ns/op
BenchmarkCopyWithLoop/10**7-4 	     100	  11964090 ns/op
BenchmarkCopyWithLoop/10**8-4 	      10	 127201115 ns/op
BenchmarkCopyWithCopyFunc/10**1-4         	50000000	        36.7 ns/op
BenchmarkCopyWithCopyFunc/10**2-4         	10000000	       147 ns/op
BenchmarkCopyWithCopyFunc/10**3-4         	 1000000	      1095 ns/op
BenchmarkCopyWithCopyFunc/10**4-4         	  200000	      7758 ns/op
BenchmarkCopyWithCopyFunc/10**5-4         	   20000	     94143 ns/op
BenchmarkCopyWithCopyFunc/10**6-4         	    2000	    952457 ns/op
BenchmarkCopyWithCopyFunc/10**7-4         	     200	   9439476 ns/op
BenchmarkCopyWithCopyFunc/10**8-4         	      20	  97646540 ns/op
PASS
ok  	github.com/keijiyoshida/goperf/slice-copy	44.070s
*/

package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkCopyWithLoop(b *testing.B) {
	benchmarkCopy(b, CopyWithLoop)
}

func BenchmarkCopyWithCopyFunc(b *testing.B) {
	benchmarkCopy(b, CopyWithCopyFunc)
}

func benchmarkCopy(b *testing.B, f func([]float64, []float64)) {
	n := 1
	for x := 1; x <= 8; x++ {
		n *= 10
		b.Run(fmt.Sprintf("10**%d", x), func(b *testing.B) {
			src := makeSlice(n)
			b.ResetTimer()
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
