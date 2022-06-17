package main

import (
	"golang.org/x/exp/constraints"
)

func Merge[T constraints.Ordered](list []T, limit int) []T {
	if len(list) < 2 {
		return list
	}
	if len(list) < limit {
		return QSort(list)
	}

	half := len(list) / 2

	out := make(chan []T)

	go func() {
		out <- Merge(list[:half], limit)
	}()
	go func() {
		out <- Merge(list[half:], limit)
	}()

	return merge(<-out, <-out)
}
