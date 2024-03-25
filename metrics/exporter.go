package exporter

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync/atomic"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	log "github.com/sirupsen/logrus"

	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
)

type Exporter struct {
	exporterConfig *config.Server
	db             *database.DB
	stopped        atomic.Bool
	tasks          tasks.Group
}

func NewExporter(ctx context.Context, db *database.DB, exporterConfig config.Server, shutdown context.CancelCauseFunc) (*Exporter, error) {

	return &Exporter{
		exporterConfig: &exporterConfig,
		db:             db,

		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in selaginella exporter: %w", err))
		}},
	}, nil
}

func (er *Exporter) Start() error {

	addr := net.JoinHostPort(er.exporterConfig.Host, strconv.Itoa(er.exporterConfig.Port))
	log.Infoln("exporter config", addr)
	log.Infoln("Starting acorus_exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
		<head><title>Acorus Exporter</title></head>
		<body>
		<h1>Acorus Exporter</h1>
		<p><a href="/metrics">Metrics</a></p>
		</body>
		</html>`))
	})

	log.Infoln("Listening on", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (er *Exporter) Stop(ctx context.Context) error {
	var result error

	if er.db != nil {
		if err := er.db.Close(); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to close DB: %w", err))
		}
	}

	er.stopped.Store(true)
	log.Info("exporter service shutdown complete")

	return nil
}

func (er *Exporter) Stopped() bool {
	return er.stopped.Load()
}
