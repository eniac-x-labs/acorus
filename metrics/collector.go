package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ChainBlockNumberMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "acorus_chain_block_number",
			Help: "The block number of any chain"},
		[]string{"chainID"},
	)
)

func init() {
	prometheus.MustRegister(ChainBlockNumberMetric)
}
