package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
)

type Exporter struct {

	//stuff1Metric prometheus.Gauge
	//stuff2Metric prometheus.Gauge

	metric1 *prometheus.Desc
	metric2 *prometheus.Desc
	metric3 *prometheus.Desc
}

func NewExporter() *Exporter {
	e := Exporter{}

	namespace := "new_metrics"
	e.metric1 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "metric1"),
		"How many metric1's.",
		[]string{"metric1string"}, prometheus.Labels{"label1":"label1val"},
	)

	e.metric2 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "metric2"),
		"How many metric2's.",
		[]string{"metric2string"}, prometheus.Labels{"label1":"label2val"},
	)

	e.metric3 = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "metric3"),
		"How many metric3's.",
		[]string{"metric3string"}, prometheus.Labels{"label1":"label3val"},
	)
	return &e
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.metric1
	ch <- e.metric2
	ch <- e.metric3
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

	a := float64(rand.Intn(100))
	ch <- prometheus.MustNewConstMetric(
		e.metric1, prometheus.GaugeValue, a, "metric1-text",
	)

	a = float64(rand.Intn(100))
	ch <- prometheus.MustNewConstMetric(
		e.metric2, prometheus.GaugeValue, a, "metric2-text",
	)

	a = float64(rand.Intn(100))
	ch <- prometheus.MustNewConstMetric(
		e.metric3, prometheus.GaugeValue, a, "metric3-text",
	)
}

func main() {

	e := NewExporter()

	//prometheus.MustRegister(e.stuff1Metric)
	//prometheus.MustRegister(e.stuff2Metric)

	prometheus.MustRegister(e)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
