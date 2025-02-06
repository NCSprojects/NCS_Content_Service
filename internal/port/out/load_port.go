package out

import "github.com/scienceMuseum/content-service/internal/domain"

type LoadPort interface {
    GetContentByID(id uint) (*domain.Content, error)
    GetAllContents() ([]*domain.Content, error)
    GetSchedulesByContentID(contentID uint) ([]*domain.ContentSchedule, error)
}
