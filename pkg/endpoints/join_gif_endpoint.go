package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"vcs.services.strawberryelk.internal/strawberryelk/gif-engine/pkg/utils"
)

type JoinGifEndpoint struct {
	Editor    *utils.GIFEditor
	ImagePath string
}

func (e *JoinGifEndpoint) Path() string {
	return "/join"
}

func (e *JoinGifEndpoint) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported.", http.StatusMethodNotAllowed)
		return
	}

	var req joinGifRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, url := range req.Urls {
		path, err := utils.DownloadUrlToFile(url, e.ImagePath)
		if err != nil {
			http.Error(w, fmt.Errorf("error downloading %s: %w", url, err).Error(), http.StatusBadRequest)
			return
		}
		err = e.Editor.LoadToBuffer(path)
		if err != nil {
			http.Error(w, fmt.Errorf("error loading %s: %w", path, err).Error(), http.StatusBadRequest)
			return
		}
	}

	data, err := e.Editor.Join()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "image/gif")
	w.Header().Set("Content-Length", strconv.Itoa(data.Len()))
	w.WriteHeader(http.StatusCreated)
	w.Write(data.Bytes())
}

type joinGifRequest struct {
	Urls []string `json:"urls"`
}
