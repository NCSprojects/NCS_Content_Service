package usecase

import (
	"mime/multipart"

	"github.com/scienceMuseum/content-service/internal/domain"
)

// ContentRegisterUseCase 인터페이스 (콘텐츠 등록 관련 비즈니스 로직)
type ContentManagementUseCase interface {
	SaveContent(content *domain.Content) error
	SaveContentWithImage(content *domain.Content, file multipart.File, fileHeader *multipart.FileHeader) error
	SaveSchedule(schedule []domain.ContentSchedule) error
	UpdateContent(content *domain.Content) error
	UpdateContentWithImage(content *domain.Content, file multipart.File, fileHeader *multipart.FileHeader) error
	ReorderContentRanks(idx []int, values []interface{})error
	DeleteContent(contentId uint) error
}
