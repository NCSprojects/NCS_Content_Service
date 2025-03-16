package out

import "github.com/scienceMuseum/content-service/internal/domain"

type SavePort interface {
    SaveContent(content *domain.Content) error
    SaveSchedule(schedule []domain.ContentSchedule) error
    UpdateContent(content *domain.Content) error
    UpdateRnk(contents []*domain.Content) error
    UpdateSchedule(content *domain.ContentSchedule) error
    DeleteContent(contentId uint ) error
    DeleteContentSchedule(ScheduleId uint ) error
}
