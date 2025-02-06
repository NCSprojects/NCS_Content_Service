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
	RegisterUseCase usecase.ContentRegisterUseCase
	FindUseCase     usecase.ContentFinderUseCase
	Router          *gin.Engine
}

// 애플리케이션 초기화
func InitializeApp() *App {
	// DB 연결
	database := config.InitDB()

	// Repository 생성
	repo := db.NewContentRepository(database)

	// Adapter 생성
	contentAdapter := adapter.NewContentAdapter(repo)

	// SavePort & LoadPort 변환
	var savePort out.SavePort = contentAdapter
	var loadPort out.LoadPort = contentAdapter

	// UseCase 생성
	registerUseCase := service.NewContentRegisterService(savePort)
	findUseCase := service.NewContentFinderService(loadPort)

	// 컨트롤러 생성
	controller := api.NewContentController(registerUseCase, findUseCase)

	// 라우터 설정
	r := gin.Default()
	r.GET("/contents/:id", controller.GetContentByID)
	r.GET("/contents", controller.GetAllContents)
	r.POST("/contents", controller.SaveContent)

	return &App{
		RegisterUseCase: registerUseCase,
		FindUseCase:     findUseCase,
		Router:          r,
	}
}

// 서버 실행
func (a *App) StartServer() {
	a.Router.Run(":3400")
}
