package service

import (
	"fmt"
	"mime/multipart"

	"github.com/scienceMuseum/content-service/internal/domain"
	"github.com/scienceMuseum/content-service/internal/port/out"
	"github.com/scienceMuseum/content-service/internal/usecase"
)

type ContentManagementService struct {
	savePort out.SavePort
	minioPort out.MinIOPort 
}

var _ usecase.ContentManagementUseCase = (*ContentManagementService)(nil)

// ContentManagementService 생성자
func NewContentManagementService(save out.SavePort,minio out.MinIOPort) *ContentManagementService {
	return &ContentManagementService{savePort: save ,minioPort :minio}
}

// 콘텐츠 저장
func (s *ContentManagementService) SaveContent(content *domain.Content) error {
	return s.savePort.SaveContent(content)
}

func (s *ContentManagementService) SaveContentWithImage(content *domain.Content, file multipart.File, fileHeader *multipart.FileHeader) error {
	imageURL := s.uploadImage(file,fileHeader)
	
	content.Photo = imageURL

	return s.savePort.SaveContent(content)
}

// 시간표 저장
func (s *ContentManagementService) SaveSchedule(schedule []domain.ContentSchedule) error {
	return s.savePort.SaveSchedule(schedule)
}

// 콘텐츠 수정
func (s *ContentManagementService) UpdateContent(content *domain.Content) error {
	return s.savePort.UpdateContent(content)
}

func (s *ContentManagementService) UpdateContentWithImage(content *domain.Content, file multipart.File, fileHeader *multipart.FileHeader) error {
	imageURL := s.uploadImage(file,fileHeader)
	
	content.Photo = imageURL

	return s.savePort.UpdateContent(content)
}

// 콘텐츠 순서 수정
func (s *ContentManagementService) ReorderContentRanks(idx []int, values []interface{}) error{
	return s.savePort.UpdateRnk(idx,"Rnk",values)
}

// 콘텐츠 삭제
func (s *ContentManagementService) DeleteContent(contentId uint) error {
	return s.savePort.DeleteContent(contentId)
}

// 이미지 업로드 처리 함수
func (s *ContentManagementService) uploadImage(file multipart.File, fileHeader *multipart.FileHeader) string {
	if file == nil {
		return "" // 파일이 없으면 빈 문자열 반환
	}

	imageURL, err := s.minioPort.UploadImage(file, fileHeader)
	if err != nil {
		fmt.Printf("이미지 업로드 실패: %v\n", err)
		return ""
	}

	return imageURL
}