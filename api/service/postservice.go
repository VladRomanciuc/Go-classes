package service

import (
	"errors"
	"math/rand"
	"strconv"

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
	i := rand.Int63()
	post.Id = strconv.FormatInt(i, 10)
	return db.AddPost(post)
}

func (*service) GetAll() ([]models.Post, error) {
	return db.GetAll()
}

func (*service) GetById(id string) (*models.Post, error) {
	/*
	_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	*/
	return db.GetById(id)
}

func (*service) DeleteById(id string) (*models.Post, error) {
	/*
	_, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	#*/

	return db.DeleteById(id)
}