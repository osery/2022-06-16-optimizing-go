package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"optimizing-go/04_otel/util"
	"optimizing-go/primes"
)

func servePrimes(c *gin.Context) {
	// TODO: Wrap in OTEL tracing span.

	var req util.PrimesRequest
	if err := c.BindQuery(&req); err != nil {
		return
	}

	result := primes.Primes(req.Max)

	c.IndentedJSON(http.StatusOK, result)
}

func main() {
	// TODO: Configure OTEL tracing.

	router := gin.Default()
	router.GET("/primes", servePrimes)
	_ = router.Run(":6061")
}
