package main

import (
	"net/http"

	"github.com/keshavsharma54126/social/internal/store"
)


type CreatePostPayload struct{
	Title string 	`json:"title"`
	Content 	string 		`json:"content"`
	Tags 	[]string	`json:tags`
}

func(app *application) createPostHandler(w http.ResponseWriter,r *http.Request){
	var payload CreatePostPayload
	if err:= readJson(w,r,&payload);err!=nil{
		writeJsonError(w,http.StatusBadRequest,err.Error())
		return
	}
	post:= &store.Post{
		Title: payload.Title,
		Content: payload.Content,
		UserID: 1,
		Tags:payload.Tags,
	}

	ctx:= r.Context()

	err:=app.store.Posts.Create(ctx,post)
	if err!=nil{
		writeJsonError(w,http.StatusInternalServerError,err.Error())
		return
	}
	err = writeJson(w,http.StatusCreated,post)
	if err!=nil{
		writeJsonError(w,http.StatusInternalServerError,err.Error())
	}
}