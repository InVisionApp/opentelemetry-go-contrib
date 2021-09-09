package datadog_test

import (
	"context"
	"time"

	"github.com/DataDog/sketches-go/ddsketch"

	"go.opentelemetry.io/contrib/exporters/metric/datadog"
	"github.com/InVisionApp/opentelemetry-go/api/global"
	"github.com/InVisionApp/opentelemetry-go/api/metric"
	export "github.com/InVisionApp/opentelemetry-go/sdk/export/metric"
	"github.com/InVisionApp/opentelemetry-go/sdk/metric/batcher/ungrouped"
	"github.com/InVisionApp/opentelemetry-go/sdk/metric/controller/push"
	"github.com/InVisionApp/opentelemetry-go/sdk/metric/selector/simple"
)

func ExampleExporter() {
	selector := simple.NewWithSketchMeasure(ddsketch.NewDefaultConfig())
	batcher := ungrouped.New(selector, export.NewDefaultLabelEncoder(), false)
	exp, err := datadog.NewExporter(datadog.Options{
		Tags: []string{"env:dev"},
	})
	if err != nil {
		panic(err)
	}
	defer exp.Close()
	pusher := push.New(batcher, exp, time.Second*10)
	defer pusher.Stop()
	pusher.Start()
	global.SetMeterProvider(pusher)
	meter := global.Meter("marwandist")
	m := metric.Must(meter).NewInt64Counter("mycounter")
	meter.RecordBatch(context.Background(), nil, m.Measurement(19))
}
