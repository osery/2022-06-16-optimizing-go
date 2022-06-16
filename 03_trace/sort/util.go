package main

import (
	"math/rand"
	"time"

	"golang.org/x/exp/constraints"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generate[T any](len int, generator func() T) []T {
	result := make([]T, len)
	for i, _ := range result {
		result[i] = generator()
	}
	return result
}

func merge[T constraints.Ordered](left []T, right []T) (result []T) {
	result = make([]T, len(left)+len(right))
	li, ri := 0, 0

	for i := 0; i < cap(result); i++ {
		if li >= len(left) ||
			(ri < len(right) && left[li] > right[ri]) {
			result[i] = right[ri]
			ri++
		} else {
			result[i] = left[li]
			li++
		}
	}
	return
}
