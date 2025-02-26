package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/scienceMuseum/content-service/internal/usecase"
	pb "github.com/scienceMuseum/content-service/proto"
)

// GRPCController 구조체
type GRPCController struct {
	pb.UnimplementedScheduleServiceServer
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