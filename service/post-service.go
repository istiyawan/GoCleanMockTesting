package service

import (
	"cleango/entity"
	"cleango/repository"
	"errors"
	"math/rand"
)

type PostService interface{
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll()([]entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository
)

func NewPostService( repos repository.PostRepository) PostService{
	repo = repos
	return &service{}
}

func (*service) Validate(post *entity.Post)error{
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}
	if post.Title == ""{
		err := errors.New("The post title is empty")
		return err
	}
	return nil
}

func (*service) Create(post *entity.Post)(*entity.Post,error){
	post.ID = rand.Int63()
	return repo.Save(post)

}
func (*service) FindAll()([]entity.Post, error){
	return repo.FindAll()
}
