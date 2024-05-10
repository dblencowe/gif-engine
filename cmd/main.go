package main

import (
	"context"
	"log"
	"vcs.services.strawberryelk.internal/strawberryelk/gif-engine/pkg/utils"

	"github.com/Netflix/go-env"
	"vcs.services.strawberryelk.internal/strawberryelk/gif-engine/pkg/database"
	"vcs.services.strawberryelk.internal/strawberryelk/gif-engine/pkg/endpoints"
	"vcs.services.strawberryelk.internal/strawberryelk/gif-engine/pkg/fallbacks"
	"vcs.services.strawberryelk.internal/strawberryelk/gif-engine/pkg/http"
)

var httpEndpoints []endpoints.Endpoint
var cfg Environment
var db database.DB

type Environment struct {
	MongoDBURI    string `env:"MONGODB_URI"`
	DataDirectory string `env:"DATA_DIR"`
	Extras        env.EnvSet
}

func init() {
	es, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	cfg.Extras = es
	log.Printf("Loaded Configuration: %+v", cfg)

	mdb, err := database.NewMongoDB(context.TODO(), cfg.MongoDBURI)
	if err != nil {
		log.Fatal(err)
	}
	db = mdb

	httpEndpoints = []endpoints.Endpoint{
		&endpoints.IndexGifEndpoint{
			DB:        db,
			ImagePath: cfg.DataDirectory,
		},
		&endpoints.GetGifEndpoint{
			DB:       db,
			Fallback: fallbacks.NewGiphyFallback(),
		},
		&endpoints.JoinGifEndpoint{
			Editor:    &utils.GIFEditor{},
			ImagePath: cfg.DataDirectory,
		},
		&endpoints.BaseEndpoint{},
	}
}

func main() {
	defer db.Stop(context.TODO())
	http.HttpServer(httpEndpoints...)
}
