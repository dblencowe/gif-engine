package endpoints

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"vcs.services.strawberryelk.internal/strawberryelk/gif-engine/pkg/database"
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
	// @todo clean this up
	if slices.Contains(formatAcceptHeader(r.Header.Get("Accept")), "image/gif") {
		w.Header().Add("Content-Type", "image/gif")
		res, err := http.Get(rst.Url())
		if err != nil || res.StatusCode != 200 {
			log.Printf("Error fetching image from URL: %v", err)
			http.Error(w, fmt.Sprintf("Unable to fetch image from %s", rst.Url()), http.StatusInternalServerError)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Printf("Error decoding body of response: %v", err)
			http.Error(w, fmt.Sprintf("Unable to fetch image from %s", rst.Url()), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Length", strconv.Itoa(len(body)))
		w.Write(body)
		return
	}

	fmt.Fprint(w, rst.Url())
}

func formatAcceptHeader(header string) []string {
	types := strings.Split(header, ",")
	for i, t := range types {
		types[i] = strings.Trim(t, " ")
	}
	return types
}
