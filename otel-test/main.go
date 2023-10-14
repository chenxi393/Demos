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

package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"os/signal"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"

	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/sdk/metric"
)

func newMetricExp(w io.Writer) (metric.Exporter, error) {
	enc := json.NewEncoder(w)
	enc.SetIndent("", " ")
	return stdoutmetric.New(
		stdoutmetric.WithEncoder(enc),
		stdoutmetric.WithoutTimestamps(),
	)
}

// newExporter returns a console exporter.
func newExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}

// newResource returns a resource describing this application.
func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("fib"),
			semconv.ServiceVersion("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	return r
}

func main() {
	l := log.New(os.Stdout, "", 0)

	// Write telemetry data to a file.
	f, err := os.Create("traces.txt")
	if err != nil {
		l.Fatal(err)
	}
	defer f.Close()

	exp, err := newExporter(f)
	if err != nil {
		l.Fatal(err)
	}

	// metrics
	metrcis_f, err := os.Create("metrics.txt")
	if err != nil {
		l.Fatal(err)
	}
	defer f.Close()
	metrci_exp, err := newMetricExp(metrcis_f)
	if err != nil {
		l.Fatal(err)
	}
	//Register the exporter with an SDK via a periodic reader.

	mp := metric.NewMeterProvider(
		metric.WithReader(
			metric.NewPeriodicReader(metrci_exp,
				metric.WithInterval(10*time.Second),
				metric.WithTimeout(2*time.Second),
			)),
		metric.WithResource(newResource()),
	)
	defer func() {
		if err := mp.Shutdown(context.Background()); err != nil {
			l.Fatal((err))
		}
	}()
	otel.SetMeterProvider(mp)

	// TracerProvider 连接了遥感数据和exporter

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(newResource()),
	)
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			l.Fatal(err)
		}
	}()
	otel.SetTracerProvider(tp)

	// 生成一个信号chan 缓冲为1
	sigCh := make(chan os.Signal, 1)
	//将中断输出给到sigCh 即可以监听中断信号
	signal.Notify(sigCh, os.Interrupt)

	// 同理 创建err chan
	errCh := make(chan error)
	//inff, _ := os.Open("1.txt")
	app := NewApp(os.Stdin, l)

	// 监听err
	go func() {
		errCh <- app.Run(context.Background())
	}()

	select {
	case <-sigCh:
		l.Println("\ngoodbye")
		return
	case err := <-errCh:
		if err != nil {
			l.Fatal(err)
		}
	}
}
