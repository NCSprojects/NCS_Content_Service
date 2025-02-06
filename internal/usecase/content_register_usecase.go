package usecase

import "github.com/scienceMuseum/content-service/internal/domain"

// ContentRegisterUseCase 인터페이스 (콘텐츠 등록 관련 비즈니스 로직)
type ContentRegisterUseCase interface {
	SaveContent(content *domain.Content) error
	SaveSchedule(schedule *domain.ContentSchedule) error
}
