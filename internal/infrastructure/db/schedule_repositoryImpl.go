package db

import (
	"github.com/scienceMuseum/content-service/internal/domain"
	"gorm.io/gorm"
)

// ScheduleRepository 구조체
type ScheduleRepositoryImpl struct {
	db *gorm.DB
}

// ScheduleRepository 생성자
func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &ScheduleRepositoryImpl{db: db}
}

// Schedule 저장
func (r *ScheduleRepositoryImpl) Create(schedules []domain.ContentSchedule) error {
	return r.db.CreateInBatches(&schedules, 5).Error
}

// 특정 Content에 속한 Schedule 조회
func (r *ScheduleRepositoryImpl) GetByContentID(contentID uint) ([]*domain.ContentSchedule, error) {
	var schedules []*domain.ContentSchedule
	err := r.db.Where("content_id = ?", contentID).Find(&schedules).Error
	return schedules, err
}

// Schedule 조회 (ID 기반) - 반환 타입을 *domain.ContentSchedule 로 수정
func (r *ScheduleRepositoryImpl) GetByID(scheduleId uint) (*domain.ContentSchedule, error) {
	var schedule domain.ContentSchedule
	if err := r.db.First(&schedule, scheduleId).Error; err != nil {
		return nil, err
	}
	return &schedule, nil
}

// Schedule 업데이트
func (r *ScheduleRepositoryImpl) Update(schedule *domain.ContentSchedule) error {
	return r.db.Save(schedule).Error
}

// Schedule 삭제
func (r *ScheduleRepositoryImpl) Delete(scheduleId uint) error {
	return r.db.Delete(&domain.ContentSchedule{}, scheduleId).Error
}