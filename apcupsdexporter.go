// Package apcupsdexporter provides the Exporter type used in the
// apcupsd_exporter Prometheus exporter.
package apcupsdexporter

import (
	"log"

	"github.com/mdlayher/apcupsd"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// namespace is the top-level namespace for this apcupsd exporter.
	namespace = "apcupsd"
)

// An Exporter is a Prometheus exporter for apcupsd metrics.
// It wraps all apcupsd metrics collectors and provides a single global
// exporter which can serve metrics.
//
// It implements the prometheus.Collector interface in order to register
// with Prometheus.
type Exporter struct {
	clientFn     ClientFunc
	nominalPower float64
}

var _ prometheus.Collector = &Exporter{}

// A ClientFunc is a function which can return an apcupsd NIS client.
// ClientFuncs are invoked on each Prometheus scrape, so that connections
// can be short-lived and less likely to time out or fail.
type ClientFunc func() (*apcupsd.Client, error)

// New creates a new Exporter which collects metrics by creating a apcupsd
// client using the input ClientFunc.
func New(fn ClientFunc, nominalPower float64) *Exporter {
	return &Exporter{
		clientFn:     fn,
		nominalPower: nominalPower,
	}
}

// Describe sends all the descriptors of the collectors included to
// the provided channel.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.withCollectors(func(cs []prometheus.Collector) {
		for _, c := range cs {
			c.Describe(ch)
		}
	})
}

// Collect sends the collected metrics from each of the collectors to
// prometheus.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.withCollectors(func(cs []prometheus.Collector) {
		for _, c := range cs {
			c.Collect(ch)
		}
	})
}

// withCollectors sets up an apcupsd client and creates a set of prometheus
// collectors.  It invokes the input closure and then cleans up after the
// closure returns.
func (e *Exporter) withCollectors(fn func(cs []prometheus.Collector)) {
	c, err := e.clientFn()
	if err != nil {
		log.Printf("[ERROR] error creating apcupsd client: %v", err)
		return
	}

	cs := []prometheus.Collector{
		NewUPSCollector(c, e.nominalPower),
	}

	fn(cs)

	_ = c.Close()
}
