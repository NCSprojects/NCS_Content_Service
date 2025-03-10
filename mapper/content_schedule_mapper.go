package mapper

import (
	"time"

	"github.com/scienceMuseum/content-service/internal/domain"
	"github.com/scienceMuseum/content-service/internal/dto"
)

// 단일 Schedule을 DTO로 변환
func ToScheduleResponseDTO(schedule *domain.ContentSchedule) dto.ScheduleResponseDTO {
	return dto.ScheduleResponseDTO{
		ID:        schedule.ID,
		ContentID: schedule.ContentID,
		StartTime: schedule.StartTime.Format(time.RFC3339),
		EndTime:   schedule.EndTime.Format(time.RFC3339),
		SeatCount: schedule.GetSeatCount(schedule.AdultCount, schedule.ChildCount),
	}
}

// 여러 개의 Schedule을 DTO로 변환
func ToScheduleResponseDTOs(schedules []*domain.ContentSchedule) []dto.ScheduleResponseDTO {
	var scheduleDTOs []dto.ScheduleResponseDTO
	for _, sch := range schedules {
		scheduleDTOs = append(scheduleDTOs, ToScheduleResponseDTO(sch))
	}
	return scheduleDTOs
}

