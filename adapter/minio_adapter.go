package adapter

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/scienceMuseum/content-service/config"
)

// MinIOAdapter 구조체
type MinIOAdapter struct {
	minioClient *minio.Client
	bucketName  string
}

// MinIOAdapter 생성자
func NewMinIOAdapter(minioClient *config.MinIOClient) *MinIOAdapter {
	return &MinIOAdapter{
		minioClient: minioClient.Client,
		bucketName:  minioClient.Bucket,
	}
}

// 이미지 업로드 메서드
func (m *MinIOAdapter) UploadImage(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	objectName := fmt.Sprintf("%d_%s", time.Now().Unix(), fileHeader.Filename)
	contentType := fileHeader.Header.Get("Content-Type")

	// MinIO에 파일 업로드
	_, err := m.minioClient.PutObject(
		context.Background(),
		m.bucketName,
		objectName,
		file,
		fileHeader.Size,
		minio.PutObjectOptions{ContentType: contentType},
	)
	if err != nil {
		return "", fmt.Errorf("MinIO 업로드 실패: %w", err)
	}

	// 업로드된 파일의 URL 반환
	imageURL := fmt.Sprintf("http://localhost:9000/%s/%s", m.bucketName, objectName)
	return imageURL, nil
}