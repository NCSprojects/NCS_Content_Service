package domain

import (
	"errors"
	"time"
)

// Content 도메인 모델
type Content struct {
    ID        uint      `json:"id"`
    CodeGrp   string    `json:"codeGrp"`
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
    Rnk       string    `json:"rnk"`
    Schedules []ContentSchedule `json:"schedules"` // 1:N 관계
}

// Content 생성자 함수
func NewContent(codeGrp, title, content, photo, mainText string,subText string ,chkDup bool,duration, totSeats int, stdDate, endDate time.Time,content_rnk string) (*Content, error) {
    if title == "" || content == "" {
        return nil, errors.New("title and content cannot be empty")
    }

    return &Content{
        CodeGrp:   codeGrp,
        Title:     title,
        Content:   content,
        Photo:     photo,
        MainText:  mainText,
        SubText:   subText,
        CreatedAt: time.Now(),
        ChkDup:    chkDup,
        Duration:  duration,
        StdDate:   stdDate,
        EndDate:   endDate,
        TotSeats:  totSeats,
        Rnk: content_rnk,
    }, nil
}

// TableName 메서드 오버라이드
func (Content) TableName() string {
    return "CONTENTS" // 대문자 테이블 이름
}

// Content 도메인 메서드
func (c *Content) IsAvailable() bool {
    now := time.Now()
    return now.After(c.StdDate) && now.Before(c.EndDate)
}

func (c *Content) TotalSeats() int {
    return c.TotSeats
}
