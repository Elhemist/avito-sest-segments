package service

import (
	repository "avito-sest-segments/internal/repository"
	segment "avito-sest-segments/models"
)

type SectionService struct {
	repo repository.Segment
}

func NewSectionService(repo repository.Segment) *SectionService {
	return &SectionService{repo: repo}
}

func (s *SectionService) CreateSegment(name segment.Segment) (int, error) {
	return s.repo.CreateSegment(name)
}
