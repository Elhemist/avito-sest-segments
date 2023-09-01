package service

import (
	repository "avito-sest-segments/internal/repository"
	segment "avito-sest-segments/models"
)

type Segment interface {
	CreateSegment(string) (int, error)
	DeleteSegment(string) error
}

type User interface {
	AddUser(user segment.User) error
	CheckUser(user segment.User) ([]string, error)
	AddSegments(user segment.SegmentsToUpdate) error
	GetPartUsers(input segment.SegmentsPartAdd) error
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
		Segment: NewSegmentService(repos.Segment),
		User:    NewUserService(repos.User),
	}
}
