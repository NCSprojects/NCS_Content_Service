package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scienceMuseum/content-service/internal/dto"
	"github.com/scienceMuseum/content-service/internal/usecase"
	"github.com/scienceMuseum/content-service/mapper"
)

// ContentController 구조체
type ContentController struct {
	RegisterUseCase usecase.ContentManagementUseCase
	FindUseCase     usecase.ContentFinderUseCase
}

// ContentController 생성자
func NewContentController(registerUseCase usecase.ContentManagementUseCase, findUseCase usecase.ContentFinderUseCase) *ContentController {
	return &ContentController{
		RegisterUseCase: registerUseCase,
		FindUseCase:     findUseCase,
	}
}

// 콘텐츠 조회 API (ID 기반)
func (cc *ContentController) GetContentByID(c *gin.Context) {
	idParam := c.Param("id")

	// 문자열을 uint로 변환
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// 콘텐츠 조회
	content, err := cc.FindUseCase.GetContentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Content not found"})
		return
	}

	// Domain → DTO 변환 후 반환
	c.JSON(http.StatusOK, mapper.ToContentResponseDTO(content))
}

// 모든 콘텐츠 조회 API
func (cc *ContentController) GetAllContents(c *gin.Context) {
	contents, err := cc.FindUseCase.GetAllContents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve contents"})
		return
	}

	// 변환된 DTO 리스트 반환
	var responseDTOs []dto.ContentResponseDTO
	for _, content := range contents {
		responseDTOs = append(responseDTOs, *mapper.ToContentResponseDTO(content))
	}

	c.JSON(http.StatusOK, responseDTOs)
}

// 콘텐츠 등록 API
func (cc *ContentController) SaveContent(c *gin.Context) {
	var req dto.ContentRequestDTO
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// DTO → Domain 변환
	content := mapper.ToContentDomain(&req)

	// 콘텐츠 저장
	err := cc.RegisterUseCase.SaveContent(content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save content"})
		return
	}

	// 변환된 응답 DTO 반환
	c.JSON(http.StatusCreated, mapper.ToContentResponseDTO(content))
}

// 콘텐츠 업데이트 API
func (cc *ContentController) UpdateContent(c *gin.Context) {
	var req dto.ContentRequestDTO
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// DTO → Domain 변환
	content := mapper.ToContentDomain(&req)

	// 콘텐츠 업데이트
	err := cc.RegisterUseCase.UpdateContent(content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update content"})
		return
	}

	// 변환된 응답 DTO 반환
	c.JSON(http.StatusOK, mapper.ToContentResponseDTO(content))
}

// 콘텐츠 삭제 API
func (cc *ContentController) DeleteContent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content ID"})
		return
	}

	// 콘텐츠 삭제 수행
	err = cc.RegisterUseCase.DeleteContent(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete content"})
		return
	}

	// 성공 응답 반환
	c.JSON(http.StatusOK, gin.H{"message": "Content deleted successfully"})
}