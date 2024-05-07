package endpoints

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"vcs.services.strawberryelk.internal/strawberryelk/gif-engine/pkg/database"
	"vcs.services.strawberryelk.internal/strawberryelk/gif-engine/pkg/fallbacks"
	"vcs.services.strawberryelk.internal/strawberryelk/gif-engine/pkg/utils"
)

type GetGifEndpoint struct {
	DB database.DB
	Fallback fallbacks.Fallback
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

	if ep.Fallback != nil {
		fallbackRst, err := ep.Fallback.Search(tags)
		if err != nil {
			log.Printf("Unable to fetch fallback: %v", err)
		} else {
			rst = fallbackRst
		}
	}

	if rst == nil {
		http.Error(w, fmt.Sprintf("0 results available for tags %v", tags), http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "image/gif")
	body, err := loadImage(rst.Url())
	if err != nil {
		log.Println("Unable to load image", err)
		http.Error(w, "Unable to load image", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	w.Write(body)
}


func loadImage(path string) ([]byte, error) {
	if strings.HasPrefix(path, "http") {
		return utils.LoadBytesFromUrl(path)
	}
	return utils.LoadBytesFromFS(path)
}