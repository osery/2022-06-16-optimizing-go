package main

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

func QSort[T constraints.Ordered](list []T) []T {
	slices.Sort(list)
	return list
}
