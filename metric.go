package otelkit

import (
	"sync"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
)

var noopMeter = sync.OnceValue(func() metric.Meter {
	return noop.NewMeterProvider().Meter("")
})

// Meter is a utility function for creating a meter. It will create
// a meter with the name matching the package of the caller with the given options.
//
// This function isn't meant to be called in a tight loop. It is intended to be
// called on initialization and the meter is supposed to be retained for future use.
func Meter(mp metric.MeterProvider, options ...metric.MeterOption) metric.Meter {
	if mp == nil {
		// Potentially consider using otel.GetTracerProvider here, but
		// we know of some issues where the default tracer provider can cause
		// memory leaks if it is never initialized so just being cautious
		// and using the noop tracer because buildkit itself doesn't use the global
		// tracer provider.
		return noopMeter()
	}

	name := packageName(1)
	return mp.Meter(name, options...)
}
