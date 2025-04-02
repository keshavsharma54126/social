package main

import (
	"net/http"
	"log"
)

func (app *application) healthCheckHandler(w http.ResponseWriter,r *http.Request) {
	data:= map[string]string{
		"status":"ok",
		"env": app.config.env,
		"version": version,
	}
	if err:= writeJson(w,http.StatusOK,data);err !=nil{
		log.Print(err.Error())
	}
}
