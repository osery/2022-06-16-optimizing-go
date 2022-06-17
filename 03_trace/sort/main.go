package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/profile"
)

func main() {
	list := generate(100_000, rand.Int31)

	before := time.Now()

	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()

	Merge(list, 1000)

	fmt.Printf("%v\n", time.Now().Sub(before))
}
