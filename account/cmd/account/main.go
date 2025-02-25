package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/muhammadali7768/go-grpc-microservices/account"
	"github.com/tinrab/retry"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config

	err := envconfig.Process("", &cfg)

	if err != nil {
		log.Fatal(err)
	}

	var r account.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) error {
		var err error
		r, err = account.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	})
	defer r.Close()
	log.Println("connected to db, Listening on port 8080....")

	s := account.NewService(r)
	err = account.ListenGRPC(s, 8080)
	if err != nil {
		log.Fatal(err)
	}
}
