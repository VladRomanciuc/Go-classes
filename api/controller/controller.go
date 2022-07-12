package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/VladRomanciuc/Go-classes/api/models"
)

var (
	postService models.PostService
	postCache models.PostCache
)

type controller struct{}

func NewPostController(service models.PostService, cache models.PostCache) models.PostController {
	postService = service
	postCache = cache
	return &controller{}
}

//get post function
func (*controller) GetAll(w http.ResponseWriter, r *http.Request) {
	//Write header with type of content "json"
	w.Header().Set("Content-type", "application/json")

	posts, err := postService.GetAll()
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ServiceError{Message: "Error getting posts from firestore"})
	}

	//if no errors the header will have status 200 and body the encoded json
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (*controller) AddPost(w http.ResponseWriter, r *http.Request) {
	//Write header with type of content "json"
	w.Header().Set("Content-type", "application/json")
	//variable post of typ Post structure
	var post models.Post
	//create new json decoder for the request body and decoding post message
	reader := json.NewDecoder(r.Body).Decode(&post)
	//error handler writes header with status and display an error message
	if reader != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ServiceError{Message: "Error Unmarshaling the request"})
		return
	}

	validator := postService.Validate(&post)
	if validator != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ServiceError{Message: validator.Error()})
		return
	}
	//add the new post to posts slice
	response, err := postService.AddPost(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ServiceError{Message: "Error adding the post"})
		return
	}
	//add new post to cache
	postCache.Set(strconv.FormatInt(post.Id, 10), &post)

	//the header will have status 200 and body the encoded json
	w.WriteHeader(http.StatusOK)
	//variable result encode post to json
	json.NewEncoder(w).Encode(response)
}

func (*controller) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	Id := strings.Split(r.URL.Path, "/")[2]
	
	var post *models.Post = postCache.Get(Id)
	if post == nil {
		post, err := postService.GetById(Id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ServiceError{Message: "No posts found!"})
			return
		}
		postCache.Set(Id, post)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(post)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(post)
	}
}
func (*controller) DeleteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	Id := strings.Split(r.URL.Path, "/")[2]
	
	var post *models.Post = postCache.Del(Id)
	if post == nil {
		post, err := postService.DeleteById(Id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ServiceError{Message: "No posts found!"})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(post)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(post)
	}
}
