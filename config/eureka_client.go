package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// EurekaClient 구조체
type EurekaClient struct {
	EurekaServerURL string
	AppName         string
	Port            int
}

// NewEurekaClient: Eureka 클라이언트 생성
func NewEurekaClient() *EurekaClient {
	eurekaServerURL := os.Getenv("EUREKA_SERVER_URL")
	if eurekaServerURL == "" {
		eurekaServerURL = "http://localhost:8761/eureka"
	}

	log.Printf("Eureka 서버 URL 설정: %s", eurekaServerURL)

	return &EurekaClient{
		EurekaServerURL: eurekaServerURL,
		AppName:         "content",
		Port:            3400,
	}
}

// Register: Eureka에 서비스 등록
func (e *EurekaClient) Register() {
	log.Println("Eureka: 서비스 등록 시도 중...")

	// 등록 요청 데이터 생성
	instance := map[string]interface{}{
		"instance": map[string]interface{}{
			"app":         e.AppName,
			"hostName":    "localhost",
			"ipAddr":      "127.0.0.1",
			"statusPageUrl": fmt.Sprintf("http://localhost:%d/info", e.Port),
			"port": map[string]interface{}{
				"$":        e.Port,
				"@enabled": true,
			},
			"vipAddress": e.AppName,
			"dataCenterInfo": map[string]interface{}{
				"@class": "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo",
				"name":   "MyOwn",
			},
		},
	}

	// JSON 변환
	jsonData, err := json.Marshal(instance)
	if err != nil {
		log.Fatalf("Eureka: JSON 변환 실패 - %v", err)
	}

	// POST 요청 보내기
	url := e.EurekaServerURL + "/apps/" + e.AppName
	log.Printf("Eureka 등록 요청 URL: %s", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Eureka: 등록 요청 생성 실패 - %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Eureka: 등록 요청 실패 - %v", err)
	}
	defer resp.Body.Close()

	// 응답 코드 확인
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Println("Eureka: 서비스 등록 성공")
	} else {
		log.Printf("Eureka: 서비스 등록 실패, 응답 코드: %d", resp.StatusCode)
	}
}

// Deregister: Eureka에서 서비스 등록 해제
func (e *EurekaClient) Deregister() {
	log.Println("Eureka: 서비스 등록 해제 시도 중...")

	url := e.EurekaServerURL + "/apps/" + e.AppName + "/localhost"
	log.Printf("Eureka 해제 요청 URL: %s", url)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatalf("Eureka: 등록 해제 요청 생성 실패 - %v", err)
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Eureka: 등록 해제 요청 실패 - %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Println("Eureka: 서비스 등록 해제 성공")
	} else {
		log.Printf("Eureka: 서비스 등록 해제 실패, 응답 코드: %d", resp.StatusCode)
	}
}