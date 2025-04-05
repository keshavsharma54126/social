package main

import (
	"net/http"
)


func(app *application)sendOtpHandler(w http.ResponseWriter,r *http.Request){
	countryCode := r.URL.Query().Get("countryCode")
	phoneNumber := r.URL.Query().Get("phoneNo")

	if phoneNumber ==""{
		writeJsonError(w,http.StatusInternalServerError,"phoneNumber is empty")
	}
	if countryCode == ""{
		writeJson(w,http.StatusInternalServerError,"countryCode is empty")
	}
	data:="message sent successfully"
	writeJson(w,http.StatusOK,data)
}

func(app *application)resendOtpHandler(w http.ResponseWriter,r *http.Request){

}

func(app *application)verityOtpHandler(w http.ResponseWriter,r *http.Request){
	
}