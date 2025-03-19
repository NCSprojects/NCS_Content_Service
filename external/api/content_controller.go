package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scienceMuseum/content-service/common"
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

// 같은 층수의 컨텐츠들 조회
func (cc *ContentController) GetContentsByFloor(c *gin.Context) {
    floor := c.Param("floor")

    contents, err := cc.FindUseCase.GetContentByFloor(floor)
    if err != nil || len(contents) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No contents found for this floor"})
        return
    }

    // 변환된 DTO 리스트 반환
    c.JSON(http.StatusOK, mapper.ToContentResponseDTOs(contents))
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

func (c *ContentController) SaveContentWithImage(ctx *gin.Context) {
	// 1. JSON DTO 데이터 가져오기
	dtoData := ctx.PostForm("dto")
	var request dto.ContentRequestDTO
	if err := json.Unmarshal([]byte(dtoData), &request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON 파싱 실패"})
		return
	}

	// DTO → Domain 변환
	content := mapper.ToContentDomain(&request)

	// 3. 이미지 파일 가져오기
	file, fileHeader, err := ctx.Request.FormFile("image")
	if err != nil {
		file = nil // 이미지가 없을 경우 nil 처리
	}

	// 4. UseCase 호출 (이미지 업로드 포함)
	err = c.RegisterUseCase.SaveContentWithImage(content, file, fileHeader)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "콘텐츠 저장 실패"})
		return
	}

	// 5. 성공 응답
	ctx.JSON(http.StatusOK, gin.H{"message": "콘텐츠 저장 완료", "content": content})
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

// 콘텐츠 순서 수정 API
func (cc *ContentController) ReorderContentRanks(ctx *gin.Context) {
	var updateList []dto.UpdateRnkDTO

	// JSON 요청 바인딩 및 검증
	if err := ctx.ShouldBindJSON(&updateList); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ID 리스트 추출
	ids, err := common.GetStructFieldValues(updateList, "ID")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID 필드 조회 실패: " + err.Error()})
		return
	}

	// Rnk 리스트 추출
	values, err := common.GetStructFieldValues(updateList, "Rnk")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Rnk 필드 조회 실패: " + err.Error()})
		return
	}

	// 서비스 호출
	err = cc.RegisterUseCase.ReorderContentRanks(common.ConvertToIntSlice(ids), values)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "순서 업데이트 실패"})
		return
	}

	// 성공 응답
	ctx.JSON(http.StatusOK, gin.H{"message": "순서 업데이트 성공"})
}


// 스케줄 조회 API (Content ID 기반)
func (cc *ContentController) GetTodaySchedulesByContentId(c *gin.Context) {
	idParam := c.Param("id")

	// 문자열을 uint로 변환
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// 스케쥴 조회
	schedules, err := cc.FindUseCase.GetTodaySchedulesByContentId(uint(id))
	if err != nil || len(schedules) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedules not found"})
		return
	}

	// Domain → DTO 변환 후 반환
	c.JSON(http.StatusOK, mapper.ToScheduleResponseDTOs(schedules))
}