package dbapi

import (
	"context"
	"log"
	"github.com/VladRomanciuc/Go-classes/api/models"

	"google.golang.org/api/option"
  
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
  )
  

type collection struct{}

const (
	projectID = "pragmatic-reviews" //refactor
	collName = "posts"
)


func NewFirestoreOps() models.DbOps{
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


//refactor

func (*collection) FindByID(id string) (*models.Post, error) {
	c := context.Background()
	client, err := firestore.NewClient(c, projectID)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	dsnap, err := client.Collection(collName).Doc(id).Get(c)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	post := &models.Post{
		Id:    dsnap.Data()["Id"].(int64),
		Title: dsnap.Data()["Title"].(string),
		Text:  dsnap.Data()["Text"].(string),
	}
	return post, nil
}

func(*collection) Delete(post *models.Post) error {
	return nil
}

