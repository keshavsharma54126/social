package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/keshavsharma54126/social/internal/store"
)


type CreatePostPayload struct{
	Title string 	`json:"title"`
	Content 	string 		`json:"content"`
	Tags 	[]string	`json:"tags"`
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

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request){
	postId:= chi.URLParam(r,"postId")
	id,err := strconv.ParseInt(postId,10,64)
	if err!=nil{
		writeJsonError(w,http.StatusInternalServerError,err.Error())
	}
	if postId== ""{
		writeJsonError(w,http.StatusBadRequest,"no postid in request params")
	}
	
	ctx:= r.Context()
	post,err:= app.store.Posts.GetById(ctx,id)
	
	if err!=nil{
		switch{
		case errors.Is(err,store.ErrNotFound):
			writeJsonError(w,http.StatusNotFound,err.Error())
		default:
			writeJsonError(w,http.StatusInternalServerError,err.Error())
		}
		return
	}
	if err:=writeJson(w,http.StatusCreated,post);err!=nil{
		writeJsonError(w,http.StatusInternalServerError,err.Error())
		return
	}
}