package app

import (
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
}

// 애플리케이션 초기화
func InitializeApp() *App {
	// DB 연결
	database := config.InitDB()

	// Repository 생성
	contentRepo := db.NewContentRepository(database)
	schduleRepo := db.NewScheduleRepository(database)

	// Adapter 생성
	contentAdapter := adapter.NewContentAdapter(contentRepo,schduleRepo)

	// SavePort & LoadPort 변환
	var savePort out.SavePort = contentAdapter
	var loadPort out.LoadPort = contentAdapter

	// UseCase 생성
	registerUseCase := service.NewContentManagementService(savePort)
	findUseCase := service.NewContentFinderService(loadPort)

	// 컨트롤러 생성
	controller := api.NewContentController(registerUseCase, findUseCase)

	// 라우터 설정
	router:= api.InitializeRouter(controller)

	return &App{
		RegisterUseCase: registerUseCase,
		FindUseCase:     findUseCase,
		Router:          router,
	}
}

// 서버 실행
func (a *App) StartServer() {
	a.Router.Run(":3400")
}
