package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"optimizing-go/04_otel/util"
	"optimizing-go/primes"
)

var tracer trace.Tracer

func servePrimes(c *gin.Context) {
	_, span := tracer.Start(c.Request.Context(), "compute-primes")
	defer span.End()

	var req util.PrimesRequest
	if err := c.BindQuery(&req); err != nil {
		return
	}

	result := primes.Primes(req.Max)

	c.IndentedJSON(http.StatusOK, result)
}

func main() {
	util.SetTracerProvider("primes-service")
	tracer = otel.Tracer("04_otel/primes")

	router := gin.Default()
	router.Use(otelgin.Middleware("http-server"))
	router.GET("/primes", servePrimes)
	_ = router.Run(":6061")
}
