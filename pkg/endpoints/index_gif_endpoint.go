package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"vcs.services.strawberryelk.internal/strawberryelk/gif-engine/pkg/database"
)

type IndexGifEndpoint struct {
	DB database.DB
}

func (ep *IndexGifEndpoint) Path() string {
	return "/write"
}

func (ep *IndexGifEndpoint) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req IndexGifRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ep.DB.Write(context.TODO(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Record: %+v", req)
}

type IndexGifRequest struct {
	Url   string
	Tags []string
}
