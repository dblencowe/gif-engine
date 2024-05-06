package http

import (
	"log"
	gohttp "net/http"

	"vcs.services.strawberryelk.internal/strawberryelk/gif-engine/pkg/endpoints"
)

var HttpListenAddr = ":5000" // @todo make this into a configurable option on a http server struct

func HttpServer(eps ...endpoints.Endpoint) error {
	for _, ep := range eps {
		gohttp.HandleFunc(ep.Path(), ep.Execute)
	}
	log.Printf("http server listening on %s", HttpListenAddr)

	return gohttp.ListenAndServe(HttpListenAddr, nil)
}
