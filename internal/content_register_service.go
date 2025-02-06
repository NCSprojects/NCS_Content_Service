package service

import (
	"github.com/scienceMuseum/content-service/internal/domain"
	"github.com/scienceMuseum/content-service/internal/port/out"
	"github.com/scienceMuseum/content-service/internal/usecase"
)

type ContentRegisterService struct {
	savePort out.SavePort 
}

var _ usecase.ContentRegisterUseCase = (*ContentRegisterService)(nil)

// ContentRegisterService 생성자
func NewContentRegisterService(save out.SavePort) *ContentRegisterService {
	return &ContentRegisterService{savePort: save}
}

// 콘텐츠 저장
func (s *ContentRegisterService) SaveContent(content *domain.Content) error {
	return s.savePort.SaveContent(content)
}

// 시간표 저장
func (s *ContentRegisterService) SaveSchedule(schedule *domain.ContentSchedule) error {
	return s.savePort.SaveSchedule(schedule)
}
