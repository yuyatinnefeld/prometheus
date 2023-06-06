package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCount *prometheus.CounterVec
	cpuUsage     *prometheus.GaugeVec
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")

    if r.URL.Path == "/" {
		requestCount.WithLabelValues("home_page").Inc()

        cpuUsage.WithLabelValues("home_page").Set(0.75)
	}
}

func metricsPage(w http.ResponseWriter, r *http.Request){
    if r.URL.Path == "/metrics" {
		requestCount.WithLabelValues("metrics").Inc()
	}

    promhttp.Handler().ServeHTTP(w, r)

}

func handleRequests() {
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
		    Name: "go_app_requests_count",
            Help: "Total App HTTP Requests count",
		},
		[]string{"endpoint"},
	)

	cpuUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "go_app_cpu_usage",
			Help: "CPU usage",
		},
		[]string{"endpoint"},
	)

	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(cpuUsage)

	server := http.NewServeMux()    
	server.HandleFunc("/", homePage)
	server.HandleFunc("/metrics", metricsPage)

	port := ":8080"
	fmt.Println("Server started on port", port)
    log.Fatal(http.ListenAndServe(port, server))
}

func main() {
    fmt.Println("Golang Rest API Starting...")
    handleRequests()
}