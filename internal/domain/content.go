package domain

import (
	"errors"
	"time"
)

// Content 도메인 모델
type Content struct {
    ID        uint      `json:"id"`
    CodeGrp   string    `json:"code_grp"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    Photo     string    `json:"photo"`
    SubTitle  string    `json:"sub_title"`
    CreatedAt time.Time `json:"created_at"`
    ChkDup    bool      `json:"chk_dup"`
    Duration  int       `json:"duration"`
    StdDate   time.Time `json:"std_date"`
    EndDate   time.Time `json:"end_date"`
    TotSeats  int       `json:"tot_seats"`
    Schedules []ContentSchedule `json:"schedules"` // 1:N 관계
}

// Content 생성자 함수
func NewContent(codeGrp, title, content, photo, subTitle string, duration, totSeats int, stdDate, endDate time.Time) (*Content, error) {
    if title == "" || content == "" {
        return nil, errors.New("title and content cannot be empty")
    }

    return &Content{
        CodeGrp:   codeGrp,
        Title:     title,
        Content:   content,
        Photo:     photo,
        SubTitle:  subTitle,
        CreatedAt: time.Now(),
        ChkDup:    false,
        Duration:  duration,
        StdDate:   stdDate,
        EndDate:   endDate,
        TotSeats:  totSeats,
    }, nil
}

// Content 도메인 메서드
func (c *Content) IsAvailable() bool {
    now := time.Now()
    return now.After(c.StdDate) && now.Before(c.EndDate)
}

func (c *Content) TotalSeats() int {
    return c.TotSeats
}
