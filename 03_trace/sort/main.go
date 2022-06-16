package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	list := generate(10_000_000, rand.Int31)

	before := time.Now()

	QSort(list)

	fmt.Printf("%v\n", time.Now().Sub(before))
}
