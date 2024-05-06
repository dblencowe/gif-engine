package endpoints

import (
	"context"
	"encoding/json"
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

	var req indexGifRequest
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
	w.WriteHeader(http.StatusCreated)
}

type indexGifRequest struct {
	Url  string
	Tags []string
}
