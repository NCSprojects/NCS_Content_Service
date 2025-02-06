package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scienceMuseum/content-service/internal/domain"
	"github.com/scienceMuseum/content-service/internal/usecase"
)

// ContentController 구조체
type ContentController struct {
	RegisterUseCase usecase.ContentRegisterUseCase
	FindUseCase     usecase.ContentFinderUseCase
}

// ContentController 생성자
func NewContentController(registerUseCase usecase.ContentRegisterUseCase, findUseCase usecase.ContentFinderUseCase) *ContentController {
	return &ContentController{
		RegisterUseCase: registerUseCase,
		FindUseCase:     findUseCase,
	}
}

// 콘텐츠 조회 API
func (cc *ContentController) GetContentByID(c *gin.Context) {
	idParam := c.Param("id")

	// 문자열을 uint로 변환
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	content, err := cc.FindUseCase.GetContentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Content not found"})
		return
	}

	c.JSON(http.StatusOK, content)
}

func (cc *ContentController) GetAllContents(c *gin.Context) {
	contents, err := cc.FindUseCase.GetAllContents()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve contents"})
		return
	}
	c.JSON(200, contents)
}

// 콘텐츠 등록 API
func (cc *ContentController) SaveContent(c *gin.Context) {
	var req domain.Content
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	err := cc.RegisterUseCase.SaveContent(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to save content"})
		return
	}
	c.JSON(201, req)
}
