package dto

import "time"

// ContentRequestDTO: 콘텐츠 생성 요청을 위한 DTO
type ContentRequestDTO struct {
    CodeGrp   string    `json:"codeGrp"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    Photo     string    `json:"photo"`
    MainText  string    `json:"mainText"`
    SubText   string    `json:"subText"`
    ChkDup    bool      `json:"chkDup"`
    Duration  int       `json:"duration"`
    StdDate   time.Time `json:"stdDate"`
    EndDate   time.Time `json:"endDate"`
    TotSeats  int       `json:"totSeats"`
    Schedules []ScheduleRequestDTO `json:"schedules"`
}

// ScheduleRequestDTO: 일정 생성 요청을 위한 DTO
type ScheduleRequestDTO struct {
    StartTime string `json:"startTime"`
    EndTime   string `json:"endTime"`
}