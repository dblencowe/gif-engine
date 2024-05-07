package fallbacks

import (
	"errors"

	"github.com/peterhellberg/giphy"
)

var ErrGiphyNoResults = errors.New("no results found for search term")

type GiphyImage struct {
	Url string
}

func NewGiphyFallback() *GiphyFallback {
	return &GiphyFallback{
		client: giphy.DefaultClient,
	}
}

type GiphyFallback struct {
	client *giphy.Client
}

func (fb *GiphyFallback) Search(terms []string) (*FallbackResult, error) {
	rst, err := fb.client.Search(terms)
	if err != nil {
		return nil, err
	}

	if len(rst.Data) == 0 {
		return nil, ErrGiphyNoResults
	}

	return &FallbackResult{
		Location: rst.Data[0].Images.Original.URL,
	}, nil
}
