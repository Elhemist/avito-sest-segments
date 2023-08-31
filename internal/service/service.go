package service

import (
	repository "avito-sest-segments/internal/repository"
	segment "avito-sest-segments/models"
)

type Segment interface {
	CreateSegment(thisSegment segment.Segment) (int, error)
}

type User interface {
	AddUser(user segment.User) (segment.User, error)
	CheckUser(user segment.User) (segment.User, error)
}
type ActiveSegment interface {
	AddActiveSegment(segment.ActiveSegment) (segment.ActiveSegment, error)
	DeleteActiveSegment(segment.ActiveSegment) error
}
type Service struct {
	Segment
	User
	ActiveSegment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Section: NewSectionService(repos.Section),
		User:    NewUserService(repos.User),
	}
}
