package usecase

import (
	"github.com/scienceMuseum/content-service/internal/domain"
)

// ContentQueryUseCase 인터페이스 (콘텐츠 조회 관련 비즈니스 로직)
type ContentFinderUseCase interface {
	GetAllContents() ([]*domain.Content, error)
	GetContentByID(id uint) (*domain.Content, error)
	GetSchedulesByContentID(contentID uint) ([]*domain.ContentSchedule, error)
	GetSchedulesIdByStartTime(startTimeStr string) ([]*string, error)
}