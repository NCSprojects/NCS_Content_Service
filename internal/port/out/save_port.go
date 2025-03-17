package out

import "github.com/scienceMuseum/content-service/internal/domain"

type SavePort interface {
    SaveContent(content *domain.Content) error
    SaveSchedule(schedule []domain.ContentSchedule) error
    UpdateContent(content *domain.Content) error
    UpdateRnk(idx []int, columnName string, values []interface{}) error
    UpdateSchedule(content *domain.ContentSchedule) error
    DeleteContent(contentId uint ) error
    DeleteContentSchedule(ScheduleId uint ) error
}
