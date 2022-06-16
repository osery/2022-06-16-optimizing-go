package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/pkg/profile"

	"optimizing-go/primes"
)

type PrimesRequest struct {
	Max int `form:"max"`
}

func serve(c *gin.Context) {
	// TODO: Trace me.

	var req PrimesRequest
	if err := c.BindQuery(&req); err != nil {
		return
	}

	var primesResult, sieveResult []int
	var wg sync.WaitGroup
	wg.Add(2)

	// Fetch primes.
	go func() {
		primesResult = primes.Primes(req.Max)
		wg.Done()
	}()

	// Fetch sieve.
	go func() {
		sieveResult = primes.Sieve(req.Max)
		wg.Done()
	}()
	wg.Wait()

	c.IndentedJSON(http.StatusOK, gin.H{
		"Primes": primesResult,
		"Sieve":  sieveResult,
	})
}

func main() {
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/both", serve)

	_ = router.Run(":6062")
}
