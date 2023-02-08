module github.com/lightstep/otel-launcher-go/pipelines

go 1.16

require (
	go.opentelemetry.io/contrib/instrumentation/host v0.32.0
	go.opentelemetry.io/contrib/instrumentation/runtime v0.32.0
	go.opentelemetry.io/contrib/propagators/b3 v1.7.0
	go.opentelemetry.io/contrib/propagators/ot v1.7.0
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.30.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.30.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.7.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.7.0
	go.opentelemetry.io/otel/metric v0.30.0
	go.opentelemetry.io/otel/sdk v1.7.0
	go.opentelemetry.io/otel/sdk/metric v0.30.0
	google.golang.org/grpc v1.53.0
)

require (
	github.com/stretchr/testify v1.8.1
	go.opentelemetry.io/proto/otlp v0.16.0
	google.golang.org/protobuf v1.28.1
)
