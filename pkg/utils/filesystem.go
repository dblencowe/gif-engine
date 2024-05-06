package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
)

func DownloadUrlToFile(url, destination string) (string, error) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	parts := strings.Split(url, ".")
	ext := parts[len(parts) - 1]
	destFileName := uuid.New()
	destinationPath := fmt.Sprintf("%s/%s.%s", destination, destFileName.String(), ext)

	err = os.WriteFile(destinationPath, body, os.ModePerm)
	if err != nil {
		return "", err
	}

	return destinationPath, nil
}

func LoadBytesFromFS(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(file)
}

func LoadBytesFromUrl(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	return body, err
}