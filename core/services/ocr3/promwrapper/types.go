package promwrapper

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type functionType string

const (
	query               functionType = "query"
	observation         functionType = "observation"
	validateObservation functionType = "validateObservation"
	outcome             functionType = "outcome"
	reports             functionType = "reports"
	shouldAccept        functionType = "shouldAccept"
	shouldTransmit      functionType = "shouldTransmit"
)

var (
	buckets = []float64{
		float64(10 * time.Millisecond),
		float64(50 * time.Millisecond),
		float64(100 * time.Millisecond),
		float64(200 * time.Millisecond),
		float64(500 * time.Millisecond),
		float64(700 * time.Millisecond),
		float64(time.Second),
		float64(2 * time.Second),
		float64(5 * time.Second),
		float64(10 * time.Second),
	}

	promOCR3ReportsGenerated = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ocr3_reporting_plugin_reports_processed",
			Help: "Tracks number of reports processed/generated within by different OCR3 functions",
		},
		[]string{"chainID", "plugin", "function"},
	)
	promOCR3Durations = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "ocr3_reporting_plugin_duration",
			Help:    "The amount of time elapsed during the OCR3 plugin's function",
			Buckets: buckets,
		},
		[]string{"chainID", "plugin", "function", "success"},
	)
)
