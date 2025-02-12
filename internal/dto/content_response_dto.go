package dto

import "time"

// ContentResponseDTO: 콘텐츠 조회 응답을 위한 DTO
type ContentResponseDTO struct {
    ID        uint      `json:"id"`
    CodeGrp   string    `json:"code_grp"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    Photo     string    `json:"photo"`
    MainText  string    `json:"mainText"`
    SubText   string    `json:"subText"`
    CreatedAt time.Time `json:"created_at"`
    ChkDup    bool      `json:"chkDup"`
    Duration  int       `json:"duration"`
    StdDate   time.Time `json:"stdDate"`
    EndDate   time.Time `json:"endDate"`
    TotSeats  int       `json:"totSeats"`
    Schedules []ScheduleResponseDTO `json:"schedules"`
}

// ScheduleResponseDTO: 일정 조회 응답을 위한 DTO
type ScheduleResponseDTO struct {
    ID        uint   `json:"id"`
    ContentID uint   `json:"contentId"`
    StartTime string `json:"startTime"`
    EndTime   string `json:"endTime"`
    SeatCount int `json:"seatCount"`
}