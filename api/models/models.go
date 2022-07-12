package models

import (
	"net/http"
)

type DbOps interface {
	AddPost(post *Post) (*Post, error)
	GetAll() ([]Post, error)
	GetById(id string) (*Post, error)
	DeleteById(id string) (*Post, error)
}

type PostService interface{
	Validate(post *Post) error
	AddPost(post *Post) (*Post, error)
	GetAll() ([]Post, error)
	GetById(id string) (*Post, error)
	DeleteById(id string) (*Post, error)
}

type PostController interface{
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	AddPost(w http.ResponseWriter, r *http.Request)
	DeleteById(w http.ResponseWriter, r *http.Request)
}

type Router interface {
	GET(url string, f func(w http.ResponseWriter, r *http.Request))
	POST(url string, f func(w http.ResponseWriter, r *http.Request))
	DELETE(url string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}

type PostCache interface {
	Set(key string, value *Post)
	Get(key string) *Post
	Del(key string) *Post
}


//The structure of data to be handled + a json mapper for encoding/decoding
type Post struct{
	Id 		int64	`json:"Id"`
	Title	string	`json:"Title"`
	Text 	string	`json:"Text"`
}

type ServiceError struct{
	Message string `json:"message"`
}

