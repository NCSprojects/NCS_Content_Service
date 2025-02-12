package db

import "github.com/scienceMuseum/content-service/internal/domain" 

type ScheduleRepository interface {
    GetByID(id uint) (*domain.ContentSchedule, error)
	GetByContentID(scheduleID uint) ([]*domain.ContentSchedule, error)
    Create(schedule []domain.ContentSchedule) error
    Update(schedule *domain.ContentSchedule) error
    Delete(scheduleId uint) error
}
