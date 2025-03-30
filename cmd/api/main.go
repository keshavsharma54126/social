package main

import (
	"log"

	"github.com/keshavsharma54126/social/internal/db"
	"github.com/keshavsharma54126/social/internal/env"
	"github.com/keshavsharma54126/social/internal/store"
)

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr: env.GetString("DB_ADDR","postgres://user:adminpassword@localhost/social?sslmod=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS",30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS",30),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME","15min"),
		},
	}

	db,err := db.New(
		cfg.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err!=nil{
		log.Panic(err)
	}


	store := store.NewStorage(db)
	app := &application{
		config: cfg,
		store:  store,
	}

	router := app.mount()

	log.Fatal(app.run(router))

}
