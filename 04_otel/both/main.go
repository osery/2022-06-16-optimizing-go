package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"optimizing-go/04_otel/util"
)

var tracer trace.Tracer
var client *resty.Client

func serveBoth(c *gin.Context) {
	var req util.PrimesRequest
	if err := c.BindQuery(&req); err != nil {
		return
	}

	var primes, sieve []int
	var wg sync.WaitGroup
	wg.Add(2)

	// Fetch primes.
	go func() {
		_, _ = client.
			R().
			SetContext(c.Request.Context()).
			SetResult(&primes).
			Get(fmt.Sprintf("http://localhost:6061/primes?max=%d", req.Max))
		wg.Done()
	}()

	// Fetch sieve.
	go func() {
		_, _ = client.R().
			SetContext(c.Request.Context()).
			SetResult(&sieve).
			Get(fmt.Sprintf("http://localhost:6062/sieve?max=%d", req.Max))
		wg.Done()
	}()

	wg.Wait()

	c.IndentedJSON(http.StatusOK, gin.H{
		"Primes": primes,
		"Sieve":  sieve,
	})
}

func main() {
	util.SetTracerProvider("both-service")
	tracer = otel.Tracer("04_otel/both")

	client = resty.NewWithClient(otelhttp.DefaultClient)

	router := gin.Default()
	router.Use(otelgin.Middleware("http-server"))
	router.GET("/both", serveBoth)
	_ = router.Run(":6060")
}
