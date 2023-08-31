package service

import (
	repository "avito-sest-segments/internal/repository"
)

type SectionService struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *SectionService {
	return &SectionService{repo: repo}
}

func (s *SectionService) CreateSegment(name string) (int, error) {
	return s.repo.CreateSegment(name)
}
func (s *SectionService) DeleteSegment(name string) error {
	return s.repo.DeleteSegment(name)
}
