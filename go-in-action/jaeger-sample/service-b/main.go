package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"log"
	"net/http"
	"time"
)

func valuesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)
	resStr := "hello from jaeger-sample-go."
	time.Sleep(2 * time.Second)
	req, _ := http.NewRequest(http.MethodGet, "http://localhost/version", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Uber-Trace-Id", r.Header.Get("Uber-Trace-Id"))

	cli := &http.Client{}
	res, _ := cli.Do(req)
	if res.StatusCode == 200 {
		body, _ := io.ReadAll(res.Body)
		resStr += string(body)
	}
	fmt.Fprintf(w, resStr)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	resStr := "jaeger-sample-go-v1.0.0"
	time.Sleep(5 * time.Second)

	fmt.Fprintf(w, resStr)
}

// TraceHandler
func TraceHandler(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		spanName := r.URL.Path
		tid := r.Header.Get("Uber-Trace-Id")
		fmt.Println("TID:", tid)
		var span opentracing.Span

		if len(tid) == 0 {
			span = opentracing.GlobalTracer().StartSpan(spanName)
			opentracing.GlobalTracer().Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		} else {
			spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
			span = opentracing.GlobalTracer().StartSpan(spanName, opentracing.ChildOf(spanCtx))
		}

		span.SetTag(string(ext.Component), spanName)
		defer span.Finish()
		handler(w, r)
	}
}

// var Tracer opentracing.Tracer

func main() {
	tracer, closer := initJaegerTracer("jaeger-sample-api-go", "jaeger-sample-agent.poneding.com:6831")
	defer closer.Close()
	opentracing.InitGlobalTracer(tracer)

	http.HandleFunc("/api/values", TraceHandler(valuesHandler))
	http.HandleFunc("/version", TraceHandler(versionHandler))
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func initJaegerTracer(serviceName, jaegerAgentAddr string) (opentracing.Tracer, io.Closer) {
	cfg := config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: jaegerAgentAddr,
		},
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))

	if err != nil {
		log.Printf("ERROR: Could not initialize jaeger-sample tracer: %s", err.Error())
	}

	return tracer, closer
}
