package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"math/big"
)

var (
	MetricsNamespace string = "syncer"
)

type AcorusMetrics interface {
	L1BlockHeight(height *big.Int)
}

type SynchronizerMetrics struct {
	l1BlockHeight prometheus.Gauge
}

func NewSynchronizerMetrics(registry *prometheus.Registry, subsystem string) *SynchronizerMetrics {
	factory := With(registry)
	return &SynchronizerMetrics{
		l1BlockHeight: factory.NewGauge(prometheus.GaugeOpts{
			Namespace: MetricsNamespace,
			Subsystem: subsystem,
			Name:      "l1_block_height",
			Help:      "the latest block height indexed into the database",
		}),
	}
}

func (sm *SynchronizerMetrics) L1BlockHeight(height *big.Int) {
	sm.l1BlockHeight.Set(float64(height.Uint64()))
}
