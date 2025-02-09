// Copyright 2020 Ipalfish, Inc.
// Copyright 2022 PingCAP, Inc.
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

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ConnGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: ModuleProxy,
			Subsystem: LabelServer,
			Name:      "connections",
			Help:      "Number of connections.",
		})

	TimeJumpBackCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: ModuleProxy,
			Subsystem: LabelMonitor,
			Name:      "time_jump_back_total",
			Help:      "Counter of system time jumps backward.",
		})

	KeepAliveCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: ModuleProxy,
			Subsystem: LabelMonitor,
			Name:      "keep_alive_total",
			Help:      "Counter of proxy keep alive.",
		})
)
