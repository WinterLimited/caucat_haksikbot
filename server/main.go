package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"server/router"
	"server/store"
)

// gin을 이용해 서버를 구축하는 코드

func main() {
	// ScyllaDB에 연결하는 코드
	// TODO: 환경 변수를 이용해, ScyllaDB에 연결할 IP, KEYSPACE 주소를 설정할 수 있도록 합니다.
	// TODO: haksik, menu, user, token 테이블을 생성 및 초기화 진행
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "yourkeyspace"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// store.Store 구조체를 생성하는 코드
	s := store.NewStore(session)

	// gin을 이용해 서버를 구축하는 코드
	// router 폴더에 정의된 Load 함수를 호출해, 서버를 구동합니다.
	r := gin.Default()
	router.Load(r, s)
	r.Run(":8080")
}
