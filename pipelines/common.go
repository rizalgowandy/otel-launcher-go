// Copyright Lightstep Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pipelines

import (
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
)

type PipelineConfig struct {
	Endpoint        string
	Insecure        bool
	Headers         map[string]string
	Resource        *resource.Resource
	ReportingPeriod string
	Propagators     []string
	MetricsOptions  []otlpmetricgrpc.Option
	TraceOptions    []otlptracegrpc.Option
}

type PipelineSetupFunc func(PipelineConfig) (func() error, error)
