package router

import (
	"github.com/gin-gonic/gin"
	"server/handler"
	"server/store"
)

func loadUserRoutes(r *gin.Engine, s *store.Store) {
	userGroup := r.Group("/user")
	userHandler := handler.NewUserHandler(s)
	{
		// @api {get} /user
		// store(임시데이터베이스)에 저장된 사용자 정보를 반환합니다.
		userGroup.GET("/", userHandler.GetUser)

		// @api {get} /user/api/:id
		// 데이터 바인딩 사용
		userGroup.GET("/:id", userHandler.GetUserByID)
	}
}
