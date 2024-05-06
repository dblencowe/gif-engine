package endpoints

import (
	"io"
	"net/http"
)

type BaseEndpoint struct{}

func (ep *BaseEndpoint) Path() string {
	return "/"
}

func (ep *BaseEndpoint) Execute(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}
