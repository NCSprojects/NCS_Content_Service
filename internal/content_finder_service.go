package service

import (
	"fmt"
	"time"

	"github.com/scienceMuseum/content-service/internal/domain"
	"github.com/scienceMuseum/content-service/internal/port/out"
	"github.com/scienceMuseum/content-service/internal/usecase"
)

// ContentFinderService 구조체 (`ContentFinderUseCase` 구현)
type ContentFinderService struct {
	loadPort out.LoadPort // 조회 인터페이스
}

var _ usecase.ContentFinderUseCase = (*ContentFinderService)(nil)

// ContentFinderService 생성자
func NewContentFinderService(load out.LoadPort) *ContentFinderService {
	return &ContentFinderService{loadPort: load}
}
// 시간 변환 함수
func parseDate(dateStr string) (time.Time, error) {
	layout := "2006-01-02" // YYYY-MM-DD 형식
	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format: expected YYYY-MM-DD, got %s", dateStr)
	}
	return parsedTime, nil
}

func (s *ContentFinderService) GetSchedulesByContentID(contentID uint) ([]*domain.ContentSchedule, error) {
	panic("unimplemented")
}

// 콘텐츠 조회 (ID 기반)
func (s *ContentFinderService) GetContentByID(id uint) (*domain.Content, error) {
	return s.loadPort.GetContentByID(id)
}

// 모든 콘텐츠 조회
func (s *ContentFinderService) GetAllContents() ([]*domain.Content, error) {
	return s.loadPort.GetAllContents()
}

func (s *ContentFinderService) GetSchedulesIdByStartTime(startTime string) ([]*string, error) {
    // 문자열을 time.Time으로 변환
    parsedTime, err := parseDate(startTime)
    if err != nil {
        return nil, fmt.Errorf("invalid startTime format: %v", err)
    }

    // 스케줄 조회
    schedules, err := s.loadPort.GetSchedulesByStartTime(parsedTime)
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve schedules: %v", err)
    }

    // ID 값을 string으로 변환하여 반환
    var scheduleIDs []*string
    for _, schedule := range schedules {
        idStr := fmt.Sprintf("%d", schedule.ID) // uint → string 변환
        scheduleIDs = append(scheduleIDs, &idStr) // *string으로 변환하여 저장
    }

    return scheduleIDs, nil
}
