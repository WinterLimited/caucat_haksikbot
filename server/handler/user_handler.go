package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/service"
	"server/store"
)

type UserHandler struct {
	store *store.Store
}

func NewUserHandler(s *store.Store) *UserHandler {
	return &UserHandler{s}
}

/**
* GET /user
* store(scylladb)에 저장된 사용자 정보를 반환합니다.
 */
func (h *UserHandler) GetUser(c *gin.Context) {
	// handler.GetUser 함수를 호출해, 사용자 정보를 조회합니다.
	users, err := h.store.UserStore.GetUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": users,
	})
}

/**
* GET /user/:id
* 요청 URL에 담긴 사용자 ID를 통해 사용자의 이름과 메뉴별 평점을 반환
* TODO: {userId}를 파싱하는 과정이 너무 복잡해서 개선방법을 알아보기 -> FIX: gin 라이브러리 사용
 */
func (h *UserHandler) GetUserByID(c *gin.Context) {
	// URL에 담긴 user_id를 바인딩합니다.
	var requestData struct {
		ID int64 `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// handler.GetUserByID 함수를 호출해, 사용자 정보를 조회합니다.
	user, err := service.FindUserByID(requestData.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
