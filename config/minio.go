package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// MinIOClient 구조체 정의
type MinIOClient struct {
	Client *minio.Client
	Bucket string
}

// MinIO 클라이언트 초기화
func NewMinIOClient() *MinIOClient {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ID")
	secretAccessKey := os.Getenv("MINIO_PW")
	bucketName := os.Getenv("MINIO_BUCKET")
	useSSL, _ := strconv.ParseBool(os.Getenv("MINIO_SSL"))

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalf("MinIO 연결 실패: %v", err)
	}

	// 컨텍스트 생성
	ctx := context.Background()

	// 버킷 존재 여부 확인
	exists, err := client.BucketExists(ctx, bucketName) // context.Background() 사용
	if err != nil {
		log.Fatalf("버킷 확인 실패: %v", err)
	}
	if !exists {
		err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{}) // context 사용
		if err != nil {
			log.Fatalf("버킷 생성 실패: %v", err)
		}
		fmt.Println("✅ MinIO 버킷 생성 완료:", bucketName)
	}

	fmt.Println("MinIO 연결 성공")
	return &MinIOClient{
		Client: client,
		Bucket: bucketName,
	}
}