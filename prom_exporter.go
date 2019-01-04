package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func addTen() float64 {
	begin += 100.00
	return begin
}

var (
	begin        = 1.00
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
	userProcessed = promauto.NewCounterFunc(prometheus.CounterOpts{
		Name: "user_processed_queries_total",
		Help: "The total number of queries per user",
	},
		addTen)
)

func main() {
	recordMetrics()
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(begin)
	})
	http.ListenAndServe(":2112", nil)
}
