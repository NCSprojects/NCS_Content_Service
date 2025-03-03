package db

import (
	"time"

	"github.com/scienceMuseum/content-service/internal/domain"
) 

type ScheduleRepository interface {
    GetByID(id uint) (*domain.ContentSchedule, error)
	GetByContentID(scheduleID uint,startDate time.Time,endDate time.Time) ([]*domain.ContentSchedule, error)
    Create(schedule []domain.ContentSchedule) error
    Update(schedule *domain.ContentSchedule) error
    Delete(scheduleId uint) error
    GetByStartTime(startTime time.Time) ([]*domain.ContentSchedule, error)
}
