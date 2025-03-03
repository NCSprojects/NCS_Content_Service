package adapter

import (
	"time"

	"github.com/scienceMuseum/content-service/internal/domain"
	"github.com/scienceMuseum/content-service/internal/infrastructure/db"
	"github.com/scienceMuseum/content-service/internal/port/out"
)

// ContentAdapter 구조체 (SavePort & LoadPort 동시 구현)
type ContentAdapter struct {
	contentRepository db.ContentRepository // Repository 주입
	scheduleRepository db.ScheduleRepository
}

var _ out.SavePort = (*ContentAdapter)(nil) 
var _ out.LoadPort = (*ContentAdapter)(nil) 
// ContentAdapter 생성자
func NewContentAdapter(contentRepo db.ContentRepository ,scheduleRepo db.ScheduleRepository ) *ContentAdapter {
	return &ContentAdapter{contentRepository: contentRepo, scheduleRepository: scheduleRepo}
	
}

func (a *ContentAdapter) GetAllContents() ([]*domain.Content, error) {
	panic("unimplemented")
}

func (a *ContentAdapter) GetContentByID(id uint) (*domain.Content, error) {
	return a.contentRepository.GetByID(id)
}

func (a *ContentAdapter) GetSchedulesByContentID(contentID uint,startDate time.Time,endDate time.Time) ([]*domain.ContentSchedule, error) {
	return a.scheduleRepository.GetByContentID(contentID,startDate,endDate)
}

func (a* ContentAdapter) GetSchedulesByStartTime(startTime time.Time) ([]*domain.ContentSchedule, error) {
	return a.scheduleRepository.GetByStartTime(startTime)
}

func (a* ContentAdapter) GetSchedulesByScheduleId(scheduleId uint) (*domain.ContentSchedule, error) {
	return a.scheduleRepository.GetByID(scheduleId)
} 

func (a *ContentAdapter) SaveContent(content *domain.Content) error {
	return a.contentRepository.Create(content)
}

func (a *ContentAdapter) SaveSchedule(schedule []domain.ContentSchedule) error {
	return a.scheduleRepository.Create(schedule)
}

func (a *ContentAdapter) UpdateContent(content *domain.Content) error {
	return a.contentRepository.Update(content)
}

func (a *ContentAdapter) UpdateSchedule(schedule *domain.ContentSchedule) error{
	return a.scheduleRepository.Update(schedule)
}

func (a *ContentAdapter) DeleteContent(contentId uint ) error {
	return a.contentRepository.Delete(contentId)
}
func (a *ContentAdapter) DeleteContentSchedule(ScheduleId uint ) error{
	return a.scheduleRepository.Delete(ScheduleId)
}



