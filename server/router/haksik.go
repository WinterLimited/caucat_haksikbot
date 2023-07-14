package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/model"
	"server/repository"
	"server/store"
	"sync"
)

func loadHaksikRoutes(r *gin.Engine) {

	// @api {get} /haksik
	// store(임시데이터베이스)에 저장된 학식 정보를 반환합니다.
	r.GET("/haksik", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"haksik": store.HaksikInfo,
		})
	})

	// @api {post} /haksik
	//데이터 바인딩 사용
	r.POST("/haksik", func(c *gin.Context) {

		// JSON 형태로 전달받은 학식 정보를 바인딩합니다.
		var requestData model.FetchData
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Fetch 함수는 네트워크 요청을 수행하는데, 이 네트워크 요청은 상대적으로 시간이 많이 걸릴 수 있습니다.
		// 따라서, Fetch 함수를 실행하는데 시간이 오래 걸린다면, 다른 요청을 처리하는데에도 시간이 오래 걸리게 됩니다.
		// 이를 방지하기 위해, Fetch 함수를 고루틴으로 실행합니다.
		var wg sync.WaitGroup
		wg.Add(1)
		go repository.Fetch(requestData.TimeOfDay, requestData.Days, requestData.IsSeoul, &wg)
		wg.Wait()

		c.JSON(http.StatusOK, gin.H{
			"haksik": store.HaksikInfo,
		})
	})
}
