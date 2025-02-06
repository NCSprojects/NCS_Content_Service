package domain

import (
	"errors"
	"time"
)

// ContentSchedule 도메인 모델
type ContentSchedule struct {
    ID        uint      `json:"id"`
    ContentID uint      `json:"content_id"` // Foreign Key
    StartTime time.Time `json:"start_time"`
    EndTime   time.Time `json:"end_time"`
    SeatCount int       `json:"seat_count"`
}

// ContentSchedule 생성자 함수
func NewContentSchedule(contentID uint, startTime, endTime time.Time, seatCount int) (*ContentSchedule, error) {
    if startTime.IsZero() || endTime.IsZero() {
        return nil, errors.New("start time and end time cannot be empty")
    }
    if endTime.Before(startTime) {
        return nil, errors.New("end time cannot be before start time")
    }
    if seatCount < 0 {
        return nil, errors.New("seat count cannot be negative")
    }

    return &ContentSchedule{
        ContentID: contentID,
        StartTime: startTime,
        EndTime:   endTime,
        SeatCount: seatCount,
    }, nil
}

// ContentSchedule 도메인 메서드

// 일정이 겹치는지 확인
func (cs *ContentSchedule) IsOverlapping(other *ContentSchedule) bool {
    return cs.StartTime.Before(other.EndTime) && cs.EndTime.After(other.StartTime)
}

// 남은 좌석 수 반환
func (cs *ContentSchedule) RemainingSeats() int {
    return cs.SeatCount
}
