package router

import (
	"github.com/gin-gonic/gin"
	"server/handler"
)

func loadHaksikRoutes(r *gin.Engine) {

	haksikGroup := r.Group("/haksik")
	{
		// GET /haksik
		// store(임시데이터베이스)에 저장된 학식 정보를 반환합니다.
		haksikGroup.GET("/", handler.GetHaksikHandler)

		// GET /haksik/api/:timeOfDay/:days/:isSeoul
		haksikGroup.GET("/:timeOfDay/:days/:isSeoul", handler.GetHaksikApiHandler)
	}

}
