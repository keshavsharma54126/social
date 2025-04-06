package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter,r *http.Request,err error){
	log.Printf("internal server error,method: %s,path: %s,error: %s",r.Method,r.URL.Path,err.Error())

	writeJsonError(w,http.StatusInternalServerError,"internal server error occured")
}

func (app * application) badRequestError(w http.ResponseWriter,r *http.Request,err error){
	log.Printf("bad request error,method: %s,path: %s,error: %s",r.Method,r.URL.Path,err.Error())

	writeJsonError(w,http.StatusBadRequest,err.Error())
}

func (app * application) notFoundResponse(w http.ResponseWriter,r *http.Request,err error){
	log.Printf("not found error,method: %s,path: %s,error: %s",r.Method,r.URL.Path,err.Error())

	writeJsonError(w,http.StatusNotFound,err.Error())
}
