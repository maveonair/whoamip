package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/maveonair/whoamip/internal/metrics"
	"github.com/maveonair/whoamip/internal/router"
)

func main() {
	portPtr := flag.Int("port", 8080, "Listening Port")
	metricsPortPtr := flag.Int("metricsPort", 9100, "Metrics Listening Port")
	flag.Parse()

	listeningAddr := fmt.Sprintf("0.0.0.0:%d", *portPtr)
	metricsListeningAddr := fmt.Sprintf("0.0.0.0:%d", *metricsPortPtr)

	go metrics.Setup(metricsListeningAddr)

	router := router.NewRouter()

	server := http.Server{
		Handler:      router,
		Addr:         listeningAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Infof("Listening on tcp://%s", listeningAddr)
	log.Fatal(server.ListenAndServe())
}
