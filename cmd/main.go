package main

import (
	"vcs.services.strawberryelk.internal/strawberryelk/gif-engine/pkg/endpoints"
	"vcs.services.strawberryelk.internal/strawberryelk/gif-engine/pkg/http"
)

var httpEndpoints []endpoints.Endpoint

func init() {
	httpEndpoints = []endpoints.Endpoint{
		&endpoints.BaseEndpoint{},
	}
}

func main() {
	http.HttpServer(httpEndpoints...)
}
