package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/izaakdale/sketchbook/internal/db"
	"github.com/izaakdale/sketchbook/internal/router"
)

func main() {
	opts, err := redis.ParseURL("redis://localhost:6379")
	if err != nil {
		log.Fatal(err)
	}

	cli := redis.NewClient(opts)
	conn, err := db.New(cli)
	if err != nil {
		log.Fatal(err)
	}
	mux := router.New(conn)

	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")),
		Handler: mux,
	}

	srv.ListenAndServe()
}
