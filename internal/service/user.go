package service

import (
	repository "avito-sest-segments/internal/repository"
	segment "avito-sest-segments/models"
	"fmt"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) AddUser(user segment.User) error {

	fmt.Println(2)
	return s.repo.AddUser(user)
}
func (s *UserService) CheckUser(user segment.User) ([]segment.Segment, error) {

	return s.repo.CheckUser(user)
}
