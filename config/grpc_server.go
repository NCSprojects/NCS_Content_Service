package config

import (
	"fmt"
	"log"
	"net"

	grpcController "github.com/scienceMuseum/content-service/external/grpc"
	"github.com/scienceMuseum/content-service/internal/usecase"
	pb "github.com/scienceMuseum/content-service/proto"
	"google.golang.org/grpc"
)

// GRPCServer 구조체
type GRPCServer struct {
	pb.UnimplementedScheduleServiceServer
	FindUseCase usecase.ContentFinderUseCase
}

// NewGRPCServer: gRPC 서버 생성
func NewGRPCServer(findUseCase usecase.ContentFinderUseCase) *GRPCServer {
	return &GRPCServer{FindUseCase: findUseCase}
}

// StartGRPCServer: gRPC 서버 실행
func (s *GRPCServer) StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50070") // gRPC 서버 포트
	if err != nil {
		log.Fatalf("❌ gRPC 서버 실행 실패: %v", err)
	}

	grpcServer := grpc.NewServer()

	// gRPC 컨트롤러 생성 및 등록
	grpcController := grpcController.NewGRPCController(s.FindUseCase)
	pb.RegisterScheduleServiceServer(grpcServer, grpcController)

	fmt.Println("🚀 gRPC server started on :50070")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ gRPC 서버 실행 실패: %v", err)
	}
}