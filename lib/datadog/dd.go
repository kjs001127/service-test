package datadog

import (
	"sync"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

const ddServiceName = "ch-app-store"

var once sync.Once

func initIfNecessary() {
	once.Do(func() {
		tracer.Start(
			tracer.WithRuntimeMetrics(),
			tracer.WithService(ddServiceName),
			tracer.WithGlobalServiceName(true),
		)
	})
}
