package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func withLoop(src []float64) []float64 {
	dst := make([]float64, len(src))
	for i, v := range src {
		dst[i] = v
	}
	return dst
}

func withCopy(src []float64) []float64 {
	dst := make([]float64, len(src))
	copy(dst, src)
	return dst
}

func withAppend(src []float64) []float64 {
	return append([]float64(nil), src...)
}

var copyFuncs = []struct {
	name string
	f    func([]float64) []float64
}{
	{"withLoop", withLoop},
	{"withCopy", withCopy},
	{"withAppend", withAppend},
}

func BenchmarkCopy(b *testing.B) {
	n := 1
	for i := 1; i <= 8; i++ {
		n *= 10
		src := newSlice(n)
		for _, copyFunc := range copyFuncs {
			b.Run(fmt.Sprintf("10**%d_%s", i, copyFunc.name), func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					dst := copyFunc.f(src)
					if src[0] != dst[0] {
						b.Errorf("src[0] => %f, dst[0] => %f", src[0], dst[0])
					}
				}
			})
		}
	}
}

func newSlice(n int) []float64 {
	s := make([]float64, n)
	for i := 0; i < n; i++ {
		s[i] = rand.Float64()
	}
	return s
}
