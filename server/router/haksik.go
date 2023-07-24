package router

import (
	"github.com/gin-gonic/gin"
	"server/handler"
	"server/store"
)

func loadHaksikRoutes(r *gin.Engine, s *store.Store) {

	// /haksik 경로에 대한 라우팅을 정의합니다.
	haksikGroup := r.Group("/haksik")
	haksikhHandler := handler.NewHaksikHandler(s)
	{
		// GET /haksik
		// store(임시데이터베이스)에 저장된 학식 정보를 반환합니다.
		haksikGroup.GET("/", haksikhHandler.GetHaksikHandler)

		// GET /haksik/api/:timeOfDay/:days/:isSeoul
		haksikGroup.GET("/:timeOfDay/:days/:isSeoul", haksikhHandler.GetHaksikApiHandler)
	}
}
