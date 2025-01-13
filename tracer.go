package otelkit

import (
	"sync"

	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

var noopTracer = sync.OnceValue(func() trace.Tracer {
	return noop.NewTracerProvider().Tracer("")
})

// Tracer is a utility function for creating a tracer. It will create
// a tracer with the name matching the package of the caller with the given options.
//
// This function isn't meant to be called in a tight loop. It is intended to be
// called on initialization and the tracer is supposed to be retained for future use.
func Tracer(tp trace.TracerProvider, options ...trace.TracerOption) trace.Tracer {
	if tp == nil {
		// Potentially consider using otel.GetTracerProvider here, but
		// we know of some issues where the default tracer provider can cause
		// memory leaks if it is never initialized so just being cautious
		// and using the noop tracer because buildkit itself doesn't use the global
		// tracer provider.
		return noopTracer()
	}

	name := packageName(1)
	return tp.Tracer(name, options...)
}
