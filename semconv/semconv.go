package semconv

import (
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

var (
	SchemaURL              = semconv.SchemaURL
	ServiceNameKey         = semconv.ServiceNameKey
	TelemetrySDKLanguageGo = semconv.TelemetrySDKLanguageGo
)

func TelemetrySDKName(val string) attribute.KeyValue {
	return semconv.TelemetrySDKName(val)
}

func TelemetrySDKVersion(val string) attribute.KeyValue {
	return semconv.TelemetrySDKVersion(val)
}
