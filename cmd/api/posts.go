package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/keshavsharma54126/social/internal/store"
)


type CreatePostPayload struct{
	Title string 	`json:"title" validate:"required,max=100"`
	Content 	string 		`json:"content" validate:"required,max=1000"`
	Tags 	[]string	`json:"tags" validate:"required"`
}


func(app *application) createPostHandler(w http.ResponseWriter,r *http.Request){
	var payload CreatePostPayload
	if err:= readJson(w,r,&payload);err!=nil{
		app.badRequestError(w,r,err)
		return
	}

	err:= Validate.Struct(payload)
	if err!=nil{
		app.badRequestError(w,r,err)
		return
	}

	post:= &store.Post{
		Title: payload.Title,
		Content: payload.Content,
		UserID: 1,
		Tags:payload.Tags,
	}

	ctx:= r.Context()

	err=app.store.Posts.Create(ctx,post)
	if err!=nil{
		app.internalServerError(w,r,err)
		return
	}
	err = writeJson(w,http.StatusCreated,post)
	if err!=nil{
		app.internalServerError(w,r,err)
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request){
	postId:= chi.URLParam(r,"postId")
	id,err := strconv.ParseInt(postId,10,64)
	if err!=nil{
		app.internalServerError(w,r,err)
	}
	if postId== ""{
		app.internalServerError(w,r,err)
		return
	}
	
	ctx:= r.Context()
	post,err:= app.store.Posts.GetById(ctx,id)
	
	if err!=nil{
		switch{
		case errors.Is(err,store.ErrNotFound):
			app.notFoundResponse(w,r,err)
		default:
			writeJsonError(w,http.StatusInternalServerError,err.Error())
		}
		return
	}
	if err:=writeJson(w,http.StatusCreated,post);err!=nil{
		app.internalServerError(w,r,err)
		return
	}
}