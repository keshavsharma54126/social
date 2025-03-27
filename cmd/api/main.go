package main

import (
	"log"

	"github.com/keshavsharma54126/social/internal/env"
	"github.com/keshavsharma54126/social/internal/store"
)

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	store := store.NewStorage(nil)
	app := &application{
		config: cfg,
		store:  store,
	}

	router := app.mount()

	log.Fatal(app.run(router))

}
