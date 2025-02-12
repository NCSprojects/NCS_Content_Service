package db

import (
	"github.com/scienceMuseum/content-service/internal/domain"
	"gorm.io/gorm"
)

// ContentRepositoryImpl 구조체 (ORM 기반 저장소)
type ContentRepositoryImpl struct {
    db *gorm.DB
}

// ContentRepositoryImpl 생성자
func NewContentRepository(db *gorm.DB) ContentRepository {
    return &ContentRepositoryImpl{db: db}
}

// 콘텐츠 저장
func (r *ContentRepositoryImpl) Create(content *domain.Content) error {
    return r.db.Create(content).Error
}

// 콘텐츠 조회 (ID 기반)
func (r *ContentRepositoryImpl) GetByID(id uint) (*domain.Content, error) {
    var content domain.Content

	err := r.db.Preload("Schedules").First(&content, id).Error
	if err != nil {
		return nil, err
	}

	return &content, nil
}

// 콘텐츠 업데이트
func (r *ContentRepositoryImpl) Update(content *domain.Content) error {
	return r.db.Model(&domain.Content{}).Where("id = ?", content.ID).Updates(content).Error
}

// 콘텐츠 삭제
func (r *ContentRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&domain.Content{}, id).Error
}


