package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define a counter metric
var (
    httpRequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "endpoint"},
    )
)

func init() {
    // Register the metric
    prometheus.MustRegister(httpRequestsTotal)
}

func main() {
    // Handle the root endpoint
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        httpRequestsTotal.With(prometheus.Labels{"method": r.Method, "endpoint": "/"}).Inc()
        w.Write([]byte("Hello, World!"))
    })

    // Expose the metrics endpoint
    http.Handle("/metrics", promhttp.Handler())

    // Start the HTTP server
    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
