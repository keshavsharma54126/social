package main

import (
	"errors"
	"net/http"
)


func(app *application)sendOtpHandler(w http.ResponseWriter,r *http.Request){
	countryCode := r.URL.Query().Get("countryCode")
	phoneNumber := r.URL.Query().Get("phoneNo")

	if phoneNumber ==""{
		app.badRequestError(w,r,errors.New("phone no is missing"))
	}
	if countryCode == ""{
		app.badRequestError(w,r,errors.New("country code is missing"))
	}
	data:="message sent successfully"
	writeJson(w,http.StatusOK,data)
}

func(app *application)resendOtpHandler(w http.ResponseWriter,r *http.Request){

}

func(app *application)verityOtpHandler(w http.ResponseWriter,r *http.Request){

}