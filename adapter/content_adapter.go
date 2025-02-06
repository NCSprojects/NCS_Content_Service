package adapter

import (
	"github.com/scienceMuseum/content-service/internal/domain"
	"github.com/scienceMuseum/content-service/internal/infrastructure/db"
	"github.com/scienceMuseum/content-service/internal/port/out"
)

// ContentAdapter 구조체 (SavePort & LoadPort 동시 구현)
type ContentAdapter struct {
	repository db.ContentRepository // Repository 주입
}

var _ out.SavePort = (*ContentAdapter)(nil) 
var _ out.LoadPort = (*ContentAdapter)(nil) 
// ContentAdapter 생성자
func NewContentAdapter(repo db.ContentRepository) *ContentAdapter {
	return &ContentAdapter{repository: repo}
}

func (a *ContentAdapter) GetAllContents() ([]*domain.Content, error) {
	panic("unimplemented")
}

func (a *ContentAdapter) GetContentByID(id uint) (*domain.Content, error) {
	return a.repository.GetByID(id)
}

func (a *ContentAdapter) GetSchedulesByContentID(contentID uint) ([]*domain.ContentSchedule, error) {
	panic("unimplemented")
}

func (a *ContentAdapter) SaveContent(content *domain.Content) error {
	return a.repository.Create(content)
}

func (a *ContentAdapter) SaveSchedule(schedule *domain.ContentSchedule) error {
	panic("unimplemented")
}


