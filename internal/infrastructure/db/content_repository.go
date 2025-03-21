package db

import "github.com/scienceMuseum/content-service/internal/domain" 

type ContentRepository interface {
    GetByID(id uint) (*domain.Content, error)
    GetByCodeGroup(codeGrp string) ([]*domain.Content, error)
    Create(content *domain.Content) error
    Update(content *domain.Content) error
    BulkColumnUpdate(idx []int, columnName string, values []interface{}) error
    Delete(contentId uint) error
    
}