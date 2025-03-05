package mapper

import (
	"time"

	"github.com/scienceMuseum/content-service/internal/domain"
	"github.com/scienceMuseum/content-service/internal/dto"
)

// DTO → Domain 변환 (요청 DTO를 도메인 객체로 변환)
func ToContentDomain(dto *dto.ContentRequestDTO) *domain.Content {
    content := &domain.Content{
        CodeGrp:   dto.CodeGrp,
        Title:     dto.Title,
        Content:   dto.Content,
        Photo:     dto.Photo,
        MainText:  dto.MainText,
        SubText:   dto.SubText,
        ChkDup:    dto.ChkDup,
        Duration:  dto.Duration,
        StdDate:   dto.StdDate,
        EndDate:   dto.EndDate,
        TotSeats:  dto.TotSeats,
    }

    // Schedules 변환
    for _, schDTO := range dto.Schedules {
        startTime, err := time.Parse(time.RFC3339, schDTO.StartTime)
        if err != nil {
            startTime = time.Time{} // 기본값 설정
        }

        endTime, err := time.Parse(time.RFC3339, schDTO.EndTime)
        if err != nil {
            endTime = time.Time{}
        }

        content.Schedules = append(content.Schedules, domain.ContentSchedule{
            StartTime: startTime,
            EndTime:   endTime,
        })
    }

    return content
}

// Domain → DTO 변환 (도메인 객체를 응답 DTO로 변환)
func ToContentResponseDTO(content *domain.Content) *dto.ContentResponseDTO {
    response := &dto.ContentResponseDTO{
        ID:        content.ID,
        CodeGrp:   content.CodeGrp,
        Title:     content.Title,
        Content:   content.Content,
        Photo:     content.Photo,
        MainText:  content.MainText,
        SubText:   content.SubText,
        CreatedAt: content.CreatedAt,
        ChkDup:    content.ChkDup,
        Duration:  content.Duration,
        StdDate:   content.StdDate,
        EndDate:   content.EndDate,
        TotSeats:  content.TotSeats,
    }

    // Schedules 변환
    for _, sch := range content.Schedules {
        response.Schedules = append(response.Schedules, dto.ScheduleResponseDTO{
            ID:        sch.ID,
            ContentID: sch.ContentID,
            StartTime: sch.StartTime.Format(time.RFC3339),
            EndTime:   sch.EndTime.Format(time.RFC3339),
            SeatCount: sch.GetSeatCount(sch.AdultCount,sch.ChildCount),
        })
    }

    return response
}

// 복수 Contents 를 DTO 리스트로 변환
func ToContentResponseDTOs(contents []*domain.Content) []dto.ContentResponseDTO {
    var responseDTOs []dto.ContentResponseDTO
    for _, content := range contents {
        responseDTOs = append(responseDTOs, *ToContentResponseDTO(content))
    }
    return responseDTOs
}