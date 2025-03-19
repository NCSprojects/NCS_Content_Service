package common

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// DecodeBase64ToFile 함수는 Base64 문자열을 디코드하여 []byte 데이터와 파일명을 반환
func DecodeBase64ToFile(base64Data, defaultFilename string) ([]byte, string, error) {
	// "data:image/jpeg;base64,..." 형태의 접두사가 있으면 제거
	if idx := strings.Index(base64Data, ","); idx != -1 {
		base64Data = base64Data[idx+1:]
	}

	// Base64 디코딩
	decodedData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return nil, "", fmt.Errorf("base64 디코드 실패: %w", err)
	}

	return decodedData, defaultFilename, nil
}