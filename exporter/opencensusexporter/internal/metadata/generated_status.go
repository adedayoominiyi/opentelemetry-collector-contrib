// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

const (
	Type             = "opencensus"
	TracesStability  = component.StabilityLevelBeta
	MetricsStability = component.StabilityLevelBeta
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("otelcol/opencensus")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("otelcol/opencensus")
}