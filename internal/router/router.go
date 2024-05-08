package router

import (
	"encoding/json"
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexHandler()).Methods(http.MethodGet)
	return router
}

func indexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		ip, err := getIpAddr(r)
		if err != nil {
			log.WithError(err).Error()
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp, err := json.Marshal(map[string]interface{}{
			"ip": ip,
		})

		if err != nil {
			log.WithError(err).Error()
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(resp)
	}
}

func getIpAddr(r *http.Request) (string, error) {
	forwarded := r.Header.Get("X-Original-Forwarded-For")
	if forwarded != "" {
		return forwarded, nil
	}

	forwarded = r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		return forwarded, nil
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	return ip, nil
}
