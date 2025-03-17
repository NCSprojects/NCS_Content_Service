package api

import (
	"github.com/gin-gonic/gin"
)

// InitializeRouter는 컨트롤러를 받아서 라우터를 설정하는 함수
func InitializeRouter(controller *ContentController) *gin.Engine {
	r := gin.Default()

	// 라우트 설정
	contentsGroup := r.Group("/contents") // 그룹 생성
	{
		contentsGroup.GET("/:id", controller.GetContentByID)   // 특정 콘텐츠 조회
		contentsGroup.GET("/", controller.GetAllContents)      // 모든 콘텐츠 조회
		contentsGroup.POST("/", controller.SaveContent)        // 콘텐츠 생성
		contentsGroup.PUT("/:id", controller.UpdateContent)    // 콘텐츠 수정
		contentsGroup.PATCH("/reorder",controller.ReorderContentRanks) // 콘텐츠 순서 변경
		contentsGroup.DELETE("/:id", controller.DeleteContent) // 콘텐츠 삭제
		contentsGroup.GET("/schedule/:id", controller.GetTodaySchedulesByContentId) // 오늘 특정 컨텐츠 스케줄 조회 
		contentsGroup.GET("/floor/:floor", controller.GetContentsByFloor) // 해당 층에 해당하는 컨텐츠 조회
	}

	return r
}
