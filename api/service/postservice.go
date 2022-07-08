package service

import (
	"errors"
	"math/rand"

	"github.com/VladRomanciuc/Go-classes/api/models"
	"github.com/VladRomanciuc/Go-classes/api/firestoreapi"
)

var (
	fireCollection firestoreapi.PostOps = firestoreapi.NewPostOpsCollection()
)

type PostService interface{
	Validate(post *models.Post) error
	AddPost(post *models.Post) (*models.Post, error)
	GetAll() ([]models.Post, error)
}

type service struct{}

//constructor
func NewPostService() PostService{
	return &service{}
}

func (*service) Validate(post *models.Post) error {
	if post == nil {
		err := errors.New("Post is empty")
		return err
	}

	if post.Title == "" {
		err := errors.New("Title is empty")
		return err
	}

	return nil
}

func (*service) AddPost(post *models.Post) (*models.Post, error) {
	post.Id = rand.Int63()
	return fireCollection.AddPost(post)
}

func (*service) GetAll() ([]models.Post, error) {
	return fireCollection.GetAll()
}