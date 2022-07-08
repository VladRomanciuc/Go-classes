package service

import (
	"errors"
	"math/rand"

	"github.com/VladRomanciuc/Go-classes/api/models"
)

var db models.DbOps

type service struct{}

//constructor
func NewPostService(dbops models.DbOps) models.PostService{
	db = dbops
	return &service{}
}

func (*service) Validate(post *models.Post) error {
	if post == nil {
		err := errors.New("post is empty")
		return err
	}

	if post.Title == "" {
		err := errors.New("title is empty")
		return err
	}

	return nil
}

func (*service) AddPost(post *models.Post) (*models.Post, error) {
	post.Id = rand.Int63()
	return db.AddPost(post)
}

func (*service) GetAll() ([]models.Post, error) {
	return db.GetAll()
}