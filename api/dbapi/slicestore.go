package dbapi

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/VladRomanciuc/Go-classes/api/models"
)

//Temp variable posts for storage
var posts []models.Post

//get post function
func GetPosts(w http.ResponseWriter, r *http.Request) {
	//Write header with type of content "json"
	w.Header().Set("Content-type", "application/json")
	//variable result encode posts to json
	result, err :=json.Marshal(posts)
	//error handler writes header with status and display an error message
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error Marshaling posts"}`))
		return
	}
	//if no errors the header will have status 200 and body the encoded json
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	//Write header with type of content "json"
	w.Header().Set("Content-type", "application/json")
	//variable post of typ Post structure
	var post models.Post
	//create new json decoder for the request body and decoding post message
	err := json.NewDecoder(r.Body).Decode(&post)
	//error handler writes header with status and display an error message
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error Unmarshaling the request"}`))
		return
	}
	//if no errors assign the next postID to the json posted of type post
	i := rand.Int63()
	post.Id = strconv.FormatInt(i, 10)
	//add the new post to posts slice
	posts = append(posts, post)
	//the header will have status 200 and body the encoded json
	w.WriteHeader(http.StatusOK)
	//variable result encode post to json
	result, err := json.Marshal(post)
	//error handler writes header with status and display an error message
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error Marshaling the post"}`))
		return
	}
	//write body with encoded json
	w.Write(result)
}