package firestore

import (
	"context"
	"log"
	"github.com/spf13/viper"
	"github.com/VladRomanciuc/Go-classes/api/views"
  
	"cloud.google.com/go/firestore"
  )
  
type PostFire interface {
	Save(post *views.Post) (*views.Post, error)
	GetAll() ([]views.Post, error)
}

type collection struct{}
const collName = "posts"


func NewPostFirestore() PostFire{
	return &collection{}
}

func getEnv(key string) string {
	viper.SetConfigFile("fireadmin.json")
	err := viper.ReadInConfig()
  
	if err != nil {
	  log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
	  log.Fatalf("Invalid type assertion")
	}
	return value
}

func (*collection) Save(post *views.Post) (*views.Post, error) {
	c := context.Background()
	client, err := firestore.NewClient(c, getEnv("project_id"))
	if err != nil {
		log.Fatal("Failed to create Firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collName).Add(c, map[string]interface{}{
		"id": post.Id,
		"title": post.Title,
		"text": post.Text,
	})
	if err != nil {
		log.Fatal("Failed adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}


func (*collection) GetAll() ([]views.Post, error) {
	c := context.Background()
	client, err := firestore.NewClient(c, getEnv("project_id"))
	if err != nil {
		log.Fatal("Failed to create Firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []views.Post

	i := client.Collection(collName).Documents(c)
	for {
		doc, err := i.Next()
		if err != nil {
			log.Fatal("Failed to return the list of posts: %v", err)
			return nil, err
		}
		post := views.Post {
			Id: doc.Data()["id"].(int64),
			Title: doc.Data()["title"].(string),
			Text: doc.Data()["text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}

