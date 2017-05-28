package main

func CopyWithLoop(dst, src []float64) {
	for i, v := range src {
		dst[i] = v
	}
}

func CopyWithCopyFunc(dst, src []float64) {
	copy(dst, src)
}
