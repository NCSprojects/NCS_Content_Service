package service

import (
	"github.com/scienceMuseum/content-service/internal/domain"
	"github.com/scienceMuseum/content-service/internal/port/out"
	"github.com/scienceMuseum/content-service/internal/usecase"
)

type ContentManagementService struct {
	savePort out.SavePort 
}

var _ usecase.ContentManagementUseCase = (*ContentManagementService)(nil)

// ContentManagementService 생성자
func NewContentManagementService(save out.SavePort) *ContentManagementService {
	return &ContentManagementService{savePort: save}
}

// 콘텐츠 저장
func (s *ContentManagementService) SaveContent(content *domain.Content) error {
	return s.savePort.SaveContent(content)
}

// 시간표 저장
func (s *ContentManagementService) SaveSchedule(schedule []domain.ContentSchedule) error {
	return s.savePort.SaveSchedule(schedule)
}

// 콘텐츠 수정
func (s *ContentManagementService) UpdateContent(content *domain.Content) error {
	return s.savePort.UpdateContent(content)
}

// 콘텐츠 순서 수정
func (s *ContentManagementService) ReorderContentRanks(idx []int, values []interface{}) error{
	return s.savePort.UpdateRnk(idx,"Rnk",values)
}

// 콘텐츠 삭제
func (s *ContentManagementService) DeleteContent(contentId uint) error {
	return s.savePort.DeleteContent(contentId)
}
