package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter,r *http.Request){
	log.Printf("%s","internal server error",r.Method)
	

}