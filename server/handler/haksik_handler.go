package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/model"
	"server/service"
	"server/store"
	"sync"
)

/**
* GET /haksik
* store(임시데이터베이스)에 저장된 학식 정보를 반환합니다.
 */
func GetHaksikHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"haksik": store.HaksikInfo,
	})
}

/*
*
* GET /haksik/api/:timeOfDay/:days/:isSeoul
// TODO: Cache를 이용해, 네트워크 요청을 최소화하는 방법을 알아보기
*/
func GetHaksikApiHandler(c *gin.Context) {

	// URL에 담긴 학식 정보를 바인딩합니다.
	var requestData model.FetchData
	if err := c.ShouldBindUri(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch 함수는 네트워크 요청을 수행하는데, 이 네트워크 요청은 상대적으로 시간이 많이 걸릴 수 있습니다.
	// 따라서, Fetch 함수를 실행하는데 시간이 오래 걸린다면, 다른 요청을 처리하는데에도 시간이 오래 걸리게 됩니다.
	// 이를 방지하기 위해, Fetch 함수를 고루틴으로 실행합니다.
	var wg sync.WaitGroup
	wg.Add(1)
	go service.Fetch(requestData.TimeOfDay, requestData.Days, requestData.IsSeoul, &wg)
	wg.Wait()

	c.JSON(http.StatusOK, gin.H{})
}
