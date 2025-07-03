package metrics

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

var (
	metricCounters   = make(map[string]prometheus.Counter)
	metricCountersMu sync.Mutex
)

// GetOrCreateCounter returns a prometheus.Counter with the given name and help.
// If a counter with that name already exists, it returns the existing one.
// Otherwise, it creates, registers, and returns a new counter.
func GetOrCreateCounter(task_name, help string) prometheus.Counter {
	metricCountersMu.Lock()
	defer metricCountersMu.Unlock()

	name := "trustyai_" + task_name
	if c, ok := metricCounters[name]; ok {
		return c
	}

	counter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: name,
		Help: help,
	})
	metrics.Registry.MustRegister(counter)
	metricCounters[name] = counter
	return counter
}

func init() {
}
