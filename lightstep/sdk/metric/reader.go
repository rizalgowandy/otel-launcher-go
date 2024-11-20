// Copyright The OpenTelemetry Authors
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

package metric // import "github.com/lightstep/otel-launcher-go/lightstep/sdk/metric"

import (
	"context"

	"github.com/lightstep/otel-launcher-go/lightstep/sdk/metric/data"
)

// Reader is the interface used between the SDK and an
// exporter.  Control flow is bi-directional through the
// Reader, since the SDK initiates ForceFlush and Shutdown
// while the Reader initiates collection.  The Register()
// method here informs the Reader that it can begin reading,
// signaling the start of bi-directional control flow.
//
// Typically, push-based exporters that are periodic will
// implement PeroidicExporter themselves and construct a
// PeriodicReader to satisfy this interface.
//
// Pull-based exporters will typically implement Register
// themselves, since they read on demand.
type Reader interface {
	// String describes this reader.
	String() string

	// Register is called when the SDK is fully
	// configured.  The Producer passed allows the
	// Reader to begin collecting metrics using its
	// Produce() method.
	Register(Producer)

	// ForceFlush is called when MeterProvider.ForceFlush() is called.
	ForceFlush(context.Context) error

	// Shutdown is called when MeterProvider.Shutdown() is called.
	Shutdown(context.Context) error
}

// Producer is the interface used to perform collection by the reader.
type Producer interface {
	// Produce returns metrics from a single collection.
	//
	// Produce may be called concurrently.
	//
	// The `in` parameter supports re-use of memory from
	// one collection to the next.  Callers that pass `in`
	// will write metrics into the same slices and structs.
	//
	// When `in` is nil, a new Metrics object is returned.
	Produce(in *data.Metrics) data.Metrics
}
