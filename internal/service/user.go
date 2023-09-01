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
	return s.repo.AddUser(user)
}
func (s *UserService) CheckUser(user segment.User) ([]segment.Segment, error) {

	return s.repo.CheckUser(user)
}

func (s *UserService) AddSegments(input segment.SegmentsToUpdate) error {
	var err error
	err = s.repo.ExistUser(input.UserId)
	if err != nil {
		return err
	}
	for _, i := range input.Delete {
		fmt.Print(i, "  ")
		segid, err := s.repo.GetSegmentByName(i)
		if err != nil {
			return err
		}
		err = s.repo.DeleteActiveSegment(segid, input.UserId)

		if err != nil {
			return err
		}
	}
	for _, i := range input.Add {
		segid, err := s.repo.GetSegmentByName(i)
		if err != nil {
			return err
		}
		err = s.repo.AddSegments(segid, input.UserId)
		if err != nil {
			return err
		}
	}
	return err
}
