package endpoints

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"vcs.services.strawberryelk.internal/strawberryelk/gif-engine/pkg/database"
	"vcs.services.strawberryelk.internal/strawberryelk/gif-engine/pkg/utils"
)

type GetGifEndpoint struct {
	DB database.DB
}

func (ep *GetGifEndpoint) Path() string {
	return "/gif"
}

func (ep *GetGifEndpoint) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tags := r.URL.Query()["tags[]"]
	if len(tags) == 0 {
		http.Error(w, "At least 1 tag must be specified", http.StatusBadRequest)
		return
	}

	rst, err := ep.DB.FindByTags(context.TODO(), tags)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if rst == nil {
		http.Error(w, fmt.Sprintf("0 results available for tags %v", tags), http.StatusNotFound)
		return
	}
	log.Printf("Accepted types: %+v", r.Header.Get("Accept"))
	w.Header().Add("Content-Type", "image/gif")
	body, err := utils.LoadBytesFromFS(rst.Url())
	if err != nil {
		log.Println("Unable to load image", err)
		http.Error(w, "Unable to load image", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	w.Write(body)
}
