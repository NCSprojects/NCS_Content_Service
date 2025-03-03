package db

import (
	"time"

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
func (r *ScheduleRepositoryImpl) GetByContentID(contentID uint, startDate time.Time, endDate time.Time) ([]*domain.ContentSchedule, error) {
	var schedules []*domain.ContentSchedule
	err := r.db.Where("content_id = ? AND start_time >= ? AND end_time <= ?", contentID , startDate, endDate).Find(&schedules).Error
	return schedules, err
}

// Schedule 조회 (ID 기반) 
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

// 특정 StartTime을 가진 Schedule 조회
func (r *ScheduleRepositoryImpl) GetByStartTime(startTime time.Time) ([]*domain.ContentSchedule, error) {
	var schedules []*domain.ContentSchedule
	// start_time 00 시 부터 23:59:59 까지
	startOfDay := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, startTime.Location())
	endOfDay := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 23, 59, 59, 999999999, startTime.Location()) // ✅ 23:59:59

	// 해당 날짜의 모든 스케줄 조회
	err := r.db.Where("start_time >= ? AND start_time <= ?", startOfDay, endOfDay).Find(&schedules).Error

	return schedules, err
}


