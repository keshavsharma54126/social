package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/keshavsharma54126/social/internal/store"
)

type application struct {
	config config
	store store.Storage
}

type config struct {
	addr string
	db		dbConfig
	env 	string
}

type dbConfig struct{
	addr string 
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string 
}

func (app *application) mount() *chi.Mux{
	r:= chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
  // Set a timeout value on the request context (ctx), that will signal
  // through ctx.Done() that the request has timed out and further
  // processing should be stopped.
  r.Use(middleware.Timeout(60 * time.Second))	

	r.Route("/v1",func(r chi.Router){
		r.Get("/health",app.healthCheckHandler)
		r.Route("/posts",func(r chi.Router){
			r.Post("/",app.createPostHandler)
			r.Route("/{postId}",func(r chi.Router){
				r.Get("/",app.getPostHandler)
			})
		})
		r.Get("/otp",app.sendOtpHandler)
		r.Post("/otp",app.verityOtpHandler)
		r.Get("/resendOtp",app.resendOtpHandler)

		r.Route("/user",func(r chi.Router){
			
		})
	})
	return r
}

func (app *application) run(mux http.Handler ) error {
	
	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: mux,
		WriteTimeout: time.Second*30,
		ReadTimeout: time.Second*10,
		IdleTimeout: time.Minute,
	}
	log.Printf("Server has started at %v", app.config.addr)
	return srv.ListenAndServe()
}
