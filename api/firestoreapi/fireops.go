package firestoreapi

import (
	"encoding/json"
	"net/http"
	"math/rand"

	"github.com/VladRomanciuc/Go-classes/api/models"
)


var (
	fireCollection PostOps = NewPostOpsCollection()
)

//get post function
func GetPosts(w http.ResponseWriter, r *http.Request) {
	//Write header with type of content "json"
	w.Header().Set("Content-type", "application/json")

	posts, err := fireCollection.GetAll()
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error getting posts from firestore"}`))
	}

	//if no errors the header will have status 200 and body the encoded json
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
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
	post.Id = rand.Int63()
	//add the new post to posts slice
	fireCollection.Save(&post)
	//the header will have status 200 and body the encoded json
	w.WriteHeader(http.StatusOK)
	//variable result encode post to json
	json.NewEncoder(w).Encode(post)
}