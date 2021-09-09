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
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/InVisionApp/opentelemetry-go/api/global"
	metricstdout "github.com/InVisionApp/opentelemetry-go/exporters/metric/stdout"
	"github.com/InVisionApp/opentelemetry-go/sdk/metric/controller/push"

	"go.opentelemetry.io/contrib/plugins/runtime"
)

func initMeter() *push.Controller {
	pusher, err := metricstdout.NewExportPipeline(metricstdout.Config{
		Quantiles:   []float64{0.5},
		PrettyPrint: true,
	}, 10*time.Second)
	if err != nil {
		log.Panicf("failed to initialize metric stdout exporter %v", err)
	}
	global.SetMeterProvider(pusher)
	return pusher
}

func main() {
	defer initMeter().Stop()

	meter := global.Meter("runtime")

	r := runtime.New(meter, time.Second)
	err := r.Start()
	if err != nil {
		panic(err)
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan

	r.Stop()
}
