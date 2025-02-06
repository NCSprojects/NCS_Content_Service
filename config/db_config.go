package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConfig 구조체 (DB 설정 정보)
type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

// 환경 변수에서 설정 로드
func LoadDBConfig() DBConfig {
	// .env 파일 로드
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	return DBConfig{
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", "password"),
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "3306"),
		Database: getEnv("DB_NAME", "content_db"),
	}
}

// 환경 변수 가져오기 (기본값 지원)
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// InitDB 함수: 설정을 기반으로 DB 연결 초기화
func InitDB() *gorm.DB {
	dbConfig := LoadDBConfig()

	// DSN (Data Source Name) 생성
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database,
	)
	log.Println("DSN :" , dsn)

	// GORM으로 MySQL 연결
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	log.Println("✅ Successfully connected to the database")
	return db
}
