/*  Copyright (c) 2022 Avesha, Inc. All rights reserved.
 *
 *  SPDX-License-Identifier: Apache-2.0
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package metrics

import (
	"net/http"
	"os"

	"github.com/kubeslice/gateway-sidecar/pkg/logger"
	"github.com/kubeslice/kubeslice-monitoring/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// create latency metrics which has to be populated when we receive latency from tunnel
var (
	sourceClusterId = os.Getenv("CLUSTER_ID")
	remoteClusterId = os.Getenv("REMOTE_CLUSTER_ID")
	remoteGatewayId = os.Getenv("REMOTE_GATEWAY_ID")
	sourceGatewayId = os.Getenv("GATEWAY_ID")
	sliceName       = os.Getenv("SLICE_NAME")
	namespace       = "kubeslice_system"
	constlabels     = prometheus.Labels{
		"slice":                   sliceName,
		"source_slice_cluster_id": sourceClusterId,
		"remote_slice_cluster_id": remoteClusterId,
		"source_gateway_id":       sourceGatewayId,
		"remote_gateway_id":       remoteGatewayId,
	}
	TunnelUP       *prometheus.GaugeVec
	LatencyMetrics *prometheus.GaugeVec
	RxRateMetrics  *prometheus.GaugeVec
	TxRateMetrics  *prometheus.GaugeVec
	log            *logger.Logger = logger.NewLogger()
)

// method to register metrics to prometheus
func StartMetricsCollector(metricCollectorPort string) {
	metricCollectorPort = ":" + metricCollectorPort
	log.Infof("Starting metric collector @ %s", metricCollectorPort)

	http.Handle("/metrics", promhttp.Handler())

	mf, err := metrics.NewMetricsFactory(prometheus.DefaultRegisterer, metrics.MetricsFactoryOptions{
		ReportingController: "gateway-sidecar",
	})

	TunnelUP = mf.NewGauge("slicegateway_tunnel_up", "Slicegateway VPN tunnel status",
		[]string{"slice", "source_gateway_id", "source_slice_cluster_id", "remote_gateway_id", "remote_slice_cluster_id"},
	).MustCurryWith(constlabels)
	LatencyMetrics = mf.NewGauge("slicegateway_tunnel_latency", "Latency between slice gateways in milliseconds",
		[]string{"slice", "source_gateway_id", "source_slice_cluster_id", "remote_gateway_id", "remote_slice_cluster_id"},
	).MustCurryWith(constlabels)
	TxRateMetrics = mf.NewGauge("slicegateway_tunnel_txrate", "Transfer rate between slice gateways in bits per second",
		[]string{"slice", "source_gateway_id", "source_slice_cluster_id", "remote_gateway_id", "remote_slice_cluster_id"},
	).MustCurryWith(constlabels)
	RxRateMetrics = mf.NewGauge("slicegateway_tunnel_rxrate", "Receive rate between slice gateways in bits per second",
		[]string{"slice", "source_gateway_id", "source_slice_cluster_id", "remote_gateway_id", "remote_slice_cluster_id"},
	).MustCurryWith(constlabels)

	if err != nil {
		log.Error("unable to initializ metrics factory")
		os.Exit(1)
	}

	err = http.ListenAndServe(metricCollectorPort, nil)
	if err != nil {
		log.Errorf("Failed to start metric collector @ %s", metricCollectorPort)
	}
	log.Info("Started Prometheus server at", metricCollectorPort)
}
