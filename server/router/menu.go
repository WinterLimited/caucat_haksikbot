package router

import (
	"github.com/gin-gonic/gin"
	"server/handler"
	"server/store"
)

func loadMenuRoutes(r *gin.Engine, s *store.Store) {
	menuGroup := r.Group("/menu")
	{
		// GET /menu
		// store(임시데이터베이스)에 저장된 메뉴 정보를 반환합니다.
		menuGroup.GET("/", handler.GetMenusHandler)

		// POST /menu/score
		menuGroup.POST("/score", handler.PostMenuHandler)

		// GET /api/menus/:name
		// TODO: handler 만들기
		//menuGroup.GET("/:name", func(c *gin.Context) {
		//
		//	// URL에 담긴 menu_name을 바인딩합니다.
		//	var requestData struct {
		//		Name string `uri:"name" binding:"required"`
		//	}
		//
		//	if err := c.ShouldBindUri(&requestData); err != nil {
		//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//		return
		//	}
		//
		//	// handler.GetMenuByName 함수를 호출해, 메뉴 정보를 조회합니다.
		//	menu, err := service.FindMenuByName(requestData.Name)
		//	if err != nil {
		//		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		//		return
		//	}
		//
		//	c.JSON(http.StatusOK, gin.H{
		//		"menu": menu,
		//	})
		//})
	}
}
