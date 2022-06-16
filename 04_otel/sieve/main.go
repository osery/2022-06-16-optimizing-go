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

func serveSieve(c *gin.Context) {
	_, span := tracer.Start(c.Request.Context(), "compute-sieve")
	defer span.End()

	var req util.PrimesRequest
	if err := c.BindQuery(&req); err != nil {
		return
	}

	result := primes.Sieve(req.Max)

	c.IndentedJSON(http.StatusOK, result)
}

func main() {
	util.SetTracerProvider("sieve-service")
	tracer = otel.Tracer("04_otel/sieve")

	router := gin.Default()
	router.Use(otelgin.Middleware("http-server"))
	router.GET("/sieve", serveSieve)
	_ = router.Run(":6062")
}
