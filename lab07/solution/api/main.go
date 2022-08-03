package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Define a histogram that will measure the time spent in the handler for "/helloworld"
	// we will define 10 buckets using prometheus.LinearBuckets (start, step, count) this
	// will give use the following buckets 0.1, 0.2, 0.3, 0.4 ...
	// Another example of measuring the time with a histogram: https://robert-scherbarth.medium.com/measure-request-duration-with-prometheus-and-golang-adc6f4ca05fe
	helloDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "hello_request_duration_seconds",
		Help:    "Histogram for the runtime of hello world handler.",
		Buckets: prometheus.LinearBuckets(0.1, 0.1, 10),
	})

	// Define a counter vec that will monitor the type of request for "/helloworld".
	// We will index this vector after a label called method. When we will index this
	// vector to increment an element the effect will be similar to:
	// 		typeCounter[method="GET"]++
	//      typeCounter[method="POST"]++
	typeCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_per_type",
			Help: "How many HTTP requests processed, partitioned by HTTP method.",
		},
		[]string{"method"},
	)

	// Define a counter that will measure the request rate for "/helloworld"
	// This will be simple counter so calling the Inc method will be equivalent
	// to helloCounter++. This will be incremented each time we get a request.
	helloCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "hello_requests_total",
		Help: "Requests received on helloworld.",
	})

	// Define an error counter for "/sum". Similar to the metric mentioned above.
	// The only difference is that we increase an "error counter" when we fail to
	// process a request (e.g. a user provides a bad input).
	sumErrCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "sum_err_total",
		Help: "Errors on sum.",
	})

	// Define a gauge to measure the uptime
	uptimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "uptime_measure",
		Help: "Measures the uptime of the service.",
	})
)

func init() {
	// Register the prometheus metrics (obs. this could have been as well-defined in main).
	prometheus.MustRegister(helloDuration)
	prometheus.MustRegister(helloCounter)
	prometheus.MustRegister(typeCounter)
	prometheus.MustRegister(sumErrCounter)
	prometheus.MustRegister(uptimeGauge)
}

func main() {
	// Add handler for prometheus we need to configure a prometheus handler.
	http.Handle("/metrics-custom", promhttp.Handler())

	// In order to monitor the uptime we're going to create a ticker that will
	// fire each 1 second. When this fires we have to increment the gauge by one
	// in order to see the uptime increasing
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				uptimeGauge.Inc()
			}
		}
	}()

	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		// Measure the time using the histogram that you have defined and a prometheus timer
		timer := prometheus.NewTimer(helloDuration)
		defer timer.ObserveDuration()

		// Incrementing the total number of requests received on "/helloworld"
		helloCounter.Inc()

		// Incrementing the number of requests received for a given method
		// typeCounter[method=r.Method]++
		typeCounter.WithLabelValues(r.Method).Inc()

		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		strA := r.URL.Query().Get("a")
		strB := r.URL.Query().Get("b")

		a, errA := strconv.Atoi(strA)
		b, errB := strconv.Atoi(strB)

		if errA != nil || errB != nil {
			// Each time a user provides a bad input and our server fails to parse it
			// we are going to increment an error counter.
			sumErrCounter.Inc()
			return
		}

		fmt.Fprintf(w, fmt.Sprintf("%d + %d = %d", a, b, a+b))
	})

	// In order to make the configuration more flexible we are going to read the listening port from
	// an environment variable using the os package.
	port := os.Getenv("LISTEN_PORT")
	fmt.Printf("Server running (port=%s)\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
