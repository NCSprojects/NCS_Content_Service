package out

import (
	"time"

	"github.com/scienceMuseum/content-service/internal/domain"
)

type LoadPort interface {
    GetContentByID(id uint) (*domain.Content, error)
    GetSchedulesByScheduleId(id uint)(*domain.ContentSchedule,error)
    GetAllContents() ([]*domain.Content, error)
    GetSchedulesByContentID(contentID uint) ([]*domain.ContentSchedule, error)
    GetSchedulesByStartTime(startTime time.Time) ([]*domain.ContentSchedule, error)
}
