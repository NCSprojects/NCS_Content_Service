package usecase

import (
	"time"

	"github.com/scienceMuseum/content-service/internal/domain"
)

// ContentUseCase 인터페이스 (콘텐츠 조회 관련 비즈니스 로직)
type ContentFinderUseCase interface {
	GetAllContents() ([]*domain.Content, error)
	GetContentByID(id uint) (*domain.Content, error)
	GetContentByFloor(floor string) ([]*domain.Content, error)
	GetSchedulesByContentID(contentID uint,startDate time.Time,endDate time.Time) ([]*domain.ContentSchedule, error)
	GetSchedulesIdByStartTime(startTimeStr string) ([]*string, error)
	GetStartTimeBySchedulesId(scheduleId string)(string , error)
	GetTodaySchedulesByContentId(contentID uint) ([]*domain.ContentSchedule, error)
}