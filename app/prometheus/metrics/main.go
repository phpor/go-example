package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {

	buckets := []float64{2, 5, 10, 25, 50, 100}
	vec := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:      "test_http_request_histogram",
			Namespace: os.Getenv("NAMESPACE"),
			Help:      "http request count",
			Buckets:   buckets,
		}, []string{"name"})

	vec2 := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:      "test_http_request_histogram",
			Namespace: os.Getenv("NAMESPACE"),
			Help:      "http request count",
			Buckets:   buckets,
		}, []string{"name"})

	prometheus.Register(vec)
	if err := prometheus.Register(vec2); err != nil {
		fmt.Println("register error: ", err)
	}

	counter1 := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "space",
		Subsystem: "subsystem",
		Name:      "requests_total",
		Help:      "Counter of service requests made.",
	}, []string{"service_type", "service_name", "method_name", "code"})

	counter2 := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "space",
		Subsystem: "subsystem",
		Name:      "requests_total",
		Help:      "Counter of service requests made.",
	}, []string{"service_type", "service_name", "method_name", "code"})

	//counter1.WithLabelValues("aa", "bb", "cc", "d").Inc()

	prometheus.Register(counter1)
	prometheus.Register(counter2)

	listen, err := net.Listen("tcp", ":9091")
	if err != nil {
		fmt.Println("listen error: ", err)
		return
	}
	paths := strings.Split(os.Args[0], "/")
	println(paths[len(paths)-1:][0])
	http.HandleFunc("/vec/inc", func(writer http.ResponseWriter, request *http.Request) {
		f, err := strconv.ParseFloat(request.URL.Query().Get("inc"), 10)
		writer.WriteHeader(200)
		writer.Header().Set("Content-Type", "text/plain")
		if err != nil {
			writer.Write([]byte("parse fail: " + err.Error()))
		}
		go func() {
			for i := 0; i <= 10000; i++ {
				counter1.WithLabelValues("aa", "bb", "cc", "d").Inc()
			}
		}()
		go func() {
			for i := 0; i <= 10000; i++ {
				counter2.WithLabelValues("aa", "bb", "cc", "d").Inc()
			}
		}()
		vec.WithLabelValues("test").Observe(f)
		vec2.WithLabelValues("test").Observe(f)
	})

	http.Handle("/metrics", promhttp.Handler())
	http.Serve(listen, nil)
}
