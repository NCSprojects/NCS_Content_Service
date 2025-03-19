package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/scienceMuseum/content-service/adapter"
	"github.com/scienceMuseum/content-service/config"
	"github.com/scienceMuseum/content-service/external/api"
	service "github.com/scienceMuseum/content-service/internal"
	"github.com/scienceMuseum/content-service/internal/infrastructure/db"
	"github.com/scienceMuseum/content-service/internal/port/out"
	"github.com/scienceMuseum/content-service/internal/usecase"
)

// App 구조체 (의존성 및 라우터 관리)
type App struct {
	RegisterUseCase usecase.ContentManagementUseCase
	FindUseCase     usecase.ContentFinderUseCase
	Router          *gin.Engine
	GRPCServer      *config.GRPCServer
	EurekaClient    *config.EurekaClient
	MinIOClient     *config.MinIOClient
}

// 애플리케이션 초기화
func InitializeApp() *App {
	// DB 연결
	database := config.InitDB()

	// MinIO 클라이언트 생성 (환경 변수에서 로드)
	minioClient := config.NewMinIOClient()

	// Repository 생성
	contentRepo := db.NewContentRepository(database)
	schduleRepo := db.NewScheduleRepository(database)

	// Adapter 생성
	contentAdapter := adapter.NewContentAdapter(contentRepo, schduleRepo)
	minioAdapter := adapter.NewMinIOAdapter(minioClient) 

	// SavePort & LoadPort 변환
	var savePort out.SavePort = contentAdapter
	var loadPort out.LoadPort = contentAdapter
	var minioPort out.MinIOPort = minioAdapter 
	

	// UseCase 생성
	registerUseCase := service.NewContentManagementService(savePort, minioPort)
	findUseCase := service.NewContentFinderService(loadPort)

	// 컨트롤러 생성
	controller := api.NewContentController(registerUseCase, findUseCase)

	// 라우터 설정
	router := api.InitializeRouter(controller)

	// gRPC 서버 생성
	grpcServer := config.NewGRPCServer(findUseCase)

	// Eureka Client 설정
	eurekaClient := config.NewEurekaClient()
	eurekaClient.Register()

	return &App{
		RegisterUseCase: registerUseCase,
		FindUseCase:     findUseCase,
		Router:          router,
		GRPCServer:      grpcServer,
		EurekaClient:    eurekaClient,
		MinIOClient:     minioClient,
	}
}

// gRPC + HTTP 서버 동시에 실행
func (a *App) StartServer() {
	go func() { // gRPC 서버 실행 (별도 goroutine)
		a.GRPCServer.StartGRPCServer()
	}()

	fmt.Println("🚀 HTTP server started on :3400")
	log.Fatal(a.Router.Run(":3400")) // HTTP 서버 실행
}