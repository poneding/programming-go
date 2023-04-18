package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"time"

	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func valuesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)
	resStr := "hello from jaeger-go."
	time.Sleep(2 * time.Second)
	req, _ := http.NewRequest(http.MethodGet, "http://localhost/version", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Uber-Trace-Id", r.Header.Get("Uber-Trace-Id"))

	cli := &http.Client{}
	res, _ := cli.Do(req)
	if res.StatusCode == 200 {
		body, _ := ioutil.ReadAll(res.Body)
		resStr += string(body)
	}
	fmt.Fprintf(w, resStr)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	resStr := "jaeger-go-v1.0.0"
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
	tracer, closer := initJaegerTracer("jaeger-api-go", "jaeger-agent.homepartners.dev:6831")
	defer closer.Close()
	opentracing.InitGlobalTracer(tracer)

	http.HandleFunc("/api/values", TraceHandler(valuesHandler))
	http.HandleFunc("/version", TraceHandler(versionHandler))
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func initJaegerTracer(serviceName, jaegerAgentAddr string) (opentracing.Tracer, io.Closer) {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: jaegerAgentAddr,
		},
	}
	tracer, closer, err := cfg.New(serviceName, config.Logger(jaeger.StdLogger))

	if err != nil {
		log.Printf("ERROR: Could not initialize jaeger tracer: %s", err.Error())
	}

	return tracer, closer
}
