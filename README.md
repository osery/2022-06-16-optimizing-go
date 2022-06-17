# Optimizing Go

Overview of Go performance analysis and profiling tools.
Presented at Pure Storage event on the 16th of June 2022.

## Render

Hosted on
`talks.godoc.org`: https://talks.godoc.org/github.com/osery/2022-06-16-optimizing-go/main.slide.

Or run locally:

```bash
go run golang.org/x/tools/cmd/present -play
```

And navigate to: http://127.0.0.1:3999/main.slide.


## Running Jaeger All-in-One Container

```shell
docker run -it --remove --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -e COLLECTOR_OTLP_ENABLED=true \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 4317:4317 \
  -p 4318:4318 \
  -p 14250:14250 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.35
```