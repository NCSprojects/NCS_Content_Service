package common

import (
	"fmt"
	"log"
	"time"
)

// getMonthStartAndEnd: 주어진 날짜의 월 첫째 날(00:00:00)과 마지막 날(23:59:59)을 반환
func GetMonthStartAndEnd(startTime string) (time.Time, time.Time, error) {
	// 입력받은 startTime을 "2006-01-02" 형식으로 변환
	parsedTime, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		log.Printf("날짜 변환 오류: %v", err)
		return time.Time{}, time.Time{}, fmt.Errorf("invalid start time format: %w", err)
	}

	// 해당 월의 첫째 날 00:00:00
	monthStart := time.Date(parsedTime.Year(), parsedTime.Month(), 1, 0, 0, 0, 0, parsedTime.Location())

	// 해당 월의 마지막 날 23:59:59
	monthEnd := time.Date(parsedTime.Year(), parsedTime.Month()+1, 0, 23, 59, 59, 0, parsedTime.Location())

	return monthStart, monthEnd, nil
}

// 시간 변환 함수
func ParseDate(dateStr string) (time.Time, error) {
	layout := "2006-01-02" // YYYY-MM-DD 형식
	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format: expected YYYY-MM-DD, got %s", dateStr)
	}
	return parsedTime, nil
}