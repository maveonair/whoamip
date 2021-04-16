package metrics

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Setup(listeningAddr string) {
	http.Handle("/metrics", promhttp.Handler())

	log.Infof("Metrics listening on tcp://%s", listeningAddr)
	log.Fatal(http.ListenAndServe(listeningAddr, nil))
}
