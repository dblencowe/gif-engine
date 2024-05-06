package endpoints

import "net/http"

type Endpoint interface {
	Path() string
	Execute(http.ResponseWriter, *http.Request) // @todo swap this out for Echo for better error handling
}
