package main

import (
	"github.com/gin-gonic/gin"
	"server/router"
)

// gin을 이용해 서버를 구축하는 코드

func main() {

	// gin을 이용해 서버를 구축하는 코드
	// router 폴더에 정의된 Load 함수를 호출해, 서버를 구동합니다.
	r := gin.Default()
	router.Load(r)
	r.Run(":8080")
}
