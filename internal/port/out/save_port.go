package out

import "github.com/scienceMuseum/content-service/internal/domain"

type SavePort interface {
    SaveContent(content *domain.Content) error
    SaveSchedule(schedule *domain.ContentSchedule) error
}
