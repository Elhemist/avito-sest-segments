package service

import (
	repository "avito-sest-segments/internal/repository"
	segment "avito-sest-segments/models"
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
func (s *UserService) CheckUser(user segment.User) ([]string, error) {
	var userSegments []string
	segIdArr, err := s.repo.CheckUser(user)
	if err != nil {
		return userSegments, err
	}
	for _, segId := range segIdArr {
		segName, err := s.repo.GetSegmentById(segId)
		if err != nil {
			return userSegments, err
		}
		userSegments = append(userSegments, segName)

	}
	return userSegments, err
}

func (s *UserService) AddSegments(input segment.SegmentsToUpdate) error {
	var err error
	err = s.repo.ExistUser(input.UserId)
	if err != nil {
		return err
	}
	for _, i := range input.Delete {
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

// GetPartUsers(per int) ([]int, error)
func (s *UserService) GetPartUsers(input segment.SegmentsPartAdd) error {
	pickedUsers, err := s.repo.GetPartUsers(input.Part)
	if err != nil {
		return err
	}
	segId, err := s.repo.GetSegmentByName(input.SegName)
	if err != nil {
		return err
	}
	for _, usr := range pickedUsers {
		err = s.repo.AddSegments(segId, usr)
		if err != nil {
			return err
		}
	}
	return err
}
