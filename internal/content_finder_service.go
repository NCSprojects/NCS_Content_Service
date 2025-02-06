package service

import (
	"github.com/scienceMuseum/content-service/internal/domain"
	"github.com/scienceMuseum/content-service/internal/port/out"
	"github.com/scienceMuseum/content-service/internal/usecase"
)

// ContentFinderService 구조체 (`ContentFinderUseCase` 구현)
type ContentFinderService struct {
	loadPort out.LoadPort // 조회 인터페이스
}

var _ usecase.ContentFinderUseCase = (*ContentFinderService)(nil)

// ContentFinderService 생성자
func NewContentFinderService(load out.LoadPort) *ContentFinderService {
	return &ContentFinderService{loadPort: load}
}

func (s *ContentFinderService) GetSchedulesByContentID(contentID uint) ([]*domain.ContentSchedule, error) {
	panic("unimplemented")
}

// 콘텐츠 조회 (ID 기반)
func (s *ContentFinderService) GetContentByID(id uint) (*domain.Content, error) {
	return s.loadPort.GetContentByID(id)
}

// 모든 콘텐츠 조회
func (s *ContentFinderService) GetAllContents() ([]*domain.Content, error) {
	return s.loadPort.GetAllContents()
}
