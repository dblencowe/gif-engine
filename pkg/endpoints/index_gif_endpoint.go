package endpoints

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"vcs.services.strawberryelk.internal/strawberryelk/gif-engine/pkg/database"
	"vcs.services.strawberryelk.internal/strawberryelk/gif-engine/pkg/utils"
)

type IndexGifEndpoint struct {
	DB        database.DB
	ImagePath string
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

	localFile, err := utils.DownloadUrlToFile(req.Filepath, ep.ImagePath)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to download image", http.StatusInternalServerError)
		return
	}
	req.Source = req.Filepath
	req.Filepath = localFile

	err = ep.DB.Write(context.TODO(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

type indexGifRequest struct {
	Filepath string `json:"url"`
	Tags     []string
	Source   string
}
