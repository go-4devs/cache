package prometheus

import (
	"time"

	metrics "github.com/prometheus/client_golang/prometheus"
	"gitoa.ru/go-4devs/cache/mw"
)

const (
	labelSet       = "label"
	labelOperation = "operation"
)

//nolint: gochecknoglobals
var (
	hitCount = metrics.NewCounterVec(
		metrics.CounterOpts{
			Name: "cache_hit_total",
			Help: "Counter of hits cache.",
		},
		[]string{labelSet},
	)
	missCount = metrics.NewCounterVec(
		metrics.CounterOpts{
			Name: "cache_miss_total",
			Help: "Counter of misses cache.",
		},
		[]string{labelSet},
	)
	expiredCount = metrics.NewCounterVec(
		metrics.CounterOpts{
			Name: "cache_expired_total",
			Help: "Counter of expired items in cache.",
		},
		[]string{labelSet},
	)
	errorsCount = metrics.NewCounterVec(
		metrics.CounterOpts{
			Name: "cache_errors_total",
			Help: "Counter of errors in cache.",
		},
		[]string{labelSet, labelOperation},
	)
	responseTime = metrics.NewHistogramVec(
		metrics.HistogramOpts{
			Name: "cache_request_duration_seconds",
			Help: "Histogram of RT for the request cache (seconds).",
		},
		[]string{labelSet, labelOperation},
	)
)

//nolint: gochecknoinits
func init() {
	metrics.MustRegister(
		hitCount,
		missCount,
		expiredCount,
		errorsCount,
		responseTime,
	)
}

var _ mw.Metrics = Metrics{}

// Metrics prometeus.
type Metrics struct{}

// Miss inc miss error cache.
func (m Metrics) Miss(label string) {
	missCount.WithLabelValues(label).Inc()
}

// Hit increment hit cache.
func (m Metrics) Hit(label string) {
	hitCount.WithLabelValues(label).Inc()
}

// Expired increment error expired.
func (m Metrics) Expired(label string) {
	expiredCount.WithLabelValues(label).Inc()
}

// Err increment base undefined error.
func (m Metrics) Err(label string, operation string) {
	errorsCount.WithLabelValues(label, operation).Inc()
}

// Observe time from start.
func (m Metrics) Observe(label string, operation string, start time.Time) {
	responseTime.WithLabelValues(label, operation).Observe(float64(time.Since(start)) / float64(time.Second))
}
