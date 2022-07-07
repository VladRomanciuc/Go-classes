package controller

import (
	"encoding/json"
	"net/http"
)

//The structure of data to be handled + a json mapper for encoding/decoding
type Post struct{
	Id 		int		`json:"id"`
	Title	string	`json:"title"`
	Text 	string	`json:"text"`
}

//Temp variable posts for storage
var (
	posts []Post
)
func init(){
	posts = []Post{{Id:1, Title: "title 1", Text: "text 1"}}
}

//get post function
func getPosts(w http.ResponseWriter, r *http.Request) {
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

func addPost(w http.ResponseWriter, r *http.Request) {
	//Write header with type of content "json"
	w.Header().Set("Content-type", "application/json")
	//variable post of typ Post structure
	var post Post
	//create new json decoder for the request body and decoding post message
	err := json.NewDecoder(r.Body).Decode(&post)
	//error handler writes header with status and display an error message
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error Unmarshaling the request"}`))
		return
	}
	//if no errors assign the next postID to the json posted of type post
	post.Id = len(posts)+1
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