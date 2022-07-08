package firestoreapi

import (
	"context"
	"log"
	"github.com/VladRomanciuc/Go-classes/api/models"

	"google.golang.org/api/option"
  
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
  )
  
type PostOps interface {
	AddPost(post *models.Post) (*models.Post, error)
	GetAll() ([]models.Post, error)
}

type collection struct{}

const collName = "posts"


func NewPostOpsCollection() PostOps{
	return &collection{}
}


func (*collection) AddPost(post *models.Post) (*models.Post, error) {
	c := context.Background()

	opt := option.WithCredentialsFile("C:\\Users\\alina\\Desktop\\Go classes\\api\\serviceAccountKey.json")
	client, err := firestore.NewClient(c, "api-go-d910c", opt)
	if err != nil {
		log.Fatal("Failed to create Firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collName).Add(c, map[string]interface{}{
		"Id": post.Id,
		"Title": post.Title,
		"Text": post.Text,
	})
	if err != nil {
		log.Fatal("Failed adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}


func (*collection) GetAll() ([]models.Post, error) {
	c := context.Background()
	
	opt := option.WithCredentialsFile("C:\\Users\\alina\\Desktop\\Go classes\\api\\serviceAccountKey.json")
	client, err := firestore.NewClient(c, "api-go-d910c", opt)
	if err != nil {
		log.Fatal("Failed to create Firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []models.Post

	coll := client.Collection(collName)
	i := coll.Documents(c)
	defer i.Stop()
	for {
		doc, err := i.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal("Failed to return the list of posts: %v", err)
			return nil, err
		}
		post := models.Post {
			Id: doc.Data()["Id"].(int64),
			Title: doc.Data()["Title"].(string),
			Text: doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}

