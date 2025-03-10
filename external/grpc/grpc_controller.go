package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/scienceMuseum/content-service/common"
	"github.com/scienceMuseum/content-service/internal/usecase"
	pb "github.com/scienceMuseum/content-service/proto"
)

// GRPCController 구조체
type GRPCController struct {
	pb.UnimplementedScheduleServiceServer
	pb.UnimplementedContentStatsServiceServer
	FindUseCase usecase.ContentFinderUseCase
}

// NewGRPCController: gRPC 컨트롤러 생성
func NewGRPCController(findUseCase usecase.ContentFinderUseCase) *GRPCController {
	return &GRPCController{FindUseCase: findUseCase}
}

// GetScheduleIdsByStartTime: gRPC 요청 처리 (스케줄 ID 조회)
func (c *GRPCController) GetScheduleIdsByStartTime(ctx context.Context, req *pb.ScheduleRequest) (*pb.ScheduleResponse, error) {
	log.Printf("gRPC 요청 수신: GetScheduleIdsByStartTime(start_time=%s)", req.StartTime)

	// UseCase를 호출하여 스케줄 ID 조회
	scheduleIDPtrs, err := c.FindUseCase.GetSchedulesIdByStartTime(req.StartTime)
	if err != nil {
		log.Printf("gRPC 요청 실패: %v", err)
		return nil, fmt.Errorf("failed to get schedule IDs: %w", err)
	}

	// []string 변환 (gRPC 응답을 위해 []*string → []string 변환)
	scheduleIDs := make([]string, len(scheduleIDPtrs))
	for i, idPtr := range scheduleIDPtrs {
		if idPtr != nil {
			scheduleIDs[i] = *idPtr
		} else {
			scheduleIDs[i] = "" // nil 포인터 방지
		}
	}

	// 응답 데이터 생성
	return &pb.ScheduleResponse{ScheduleIds: scheduleIDs}, nil
}

func (c *GRPCController) GetStartTimeBySchedulesId(ctx context.Context, req *pb.ScheduleIdRequest) (*pb.StartTimeResponse, error) {
	log.Printf("gRPC 요청 수신: GetStartTimeBySchedulesId(schedule_id=%s)", req.ScheduleId)

	// UseCase를 호출하여 스케줄 시작 시간 조회
	startTime, err := c.FindUseCase.GetStartTimeBySchedulesId(req.ScheduleId)
	if err != nil {
		log.Printf("gRPC 요청 실패: %v", err)
		return nil, fmt.Errorf("failed to get schedule start time: %w", err)
	}

	// 응답 데이터 생성
	return &pb.StartTimeResponse{StartTime: startTime}, nil
}

// GetContentStats: gRPC 요청 처리 (콘텐츠 통계 조회)
func (c *GRPCController) GetContentStats(ctx context.Context, req *pb.ContentStatsRequest) (*pb.ContentStatsResponseList, error) {
	log.Printf("gRPC 요청 수신: GetContentStats(start_time=%s, content_id=%d)", req.StartTime, req.ContentId)

	startDate, endDate, err := common.GetMonthStartAndEnd(req.StartTime)
	if err != nil {
		log.Printf("❌ 날짜 변환 오류: %v", err)
		return nil, fmt.Errorf("invalid start time format: %w", err)
	}

	stats, err := c.FindUseCase.GetSchedulesByContentID(uint(req.ContentId), startDate, endDate)
	if err != nil {
		log.Printf("gRPC 요청 실패: %v", err)
		return nil, fmt.Errorf("failed to get content stats: %w", err)
	}

	// 응답 데이터 변환
	var responses []*pb.ContentStatsResponse
	for _, stat := range stats {
		responses = append(responses, &pb.ContentStatsResponse{
			StartTime: stat.StartTime.Format("2006-01-02 15:04:05"),
			AdCnt:     int32(stat.AdultCount),
			CdCnt:     int32(stat.ChildCount),
		})
	}

	// 응답 데이터 생성
	return &pb.ContentStatsResponseList{Responses: responses}, nil
}
