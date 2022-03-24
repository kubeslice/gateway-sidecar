package metrics

// Records latency for this sidecar
func RecordLatencyMetric(latency float64) {
	// Set Latency Gauge in prometheus
	LatencyMetrics.Set(latency)
}

// Records rx bytes for this sidecar
func RecordRxRateMetric(rxRate float64) {
	// Set Rx bytes in prometheus
	RxRateMetrics.Set(rxRate)
}

// Records tx bytes for this sidecar
func RecordTxRateMetric(txRate float64) {
	// Set Tx bytes in prometheus
	TxRateMetrics.Set(txRate)
}
