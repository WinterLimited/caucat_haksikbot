package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/service"
)

/**
* GET /menu
* 모든 메뉴에 대한 정보를 반환
 */
func GetMenusHandler(c *gin.Context) {
	menus, err := service.FindMenus()
	if err != nil {
		// 에러 처리
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"menus": menus,
	})
}

/**
* POST menu/score
* 요청 바디에 담긴 메뉴 이름과 사용자 ID를 통해 메뉴에 대한 평점을 삽입, 갱신
* 요청 body 예시
*	{
*		"menuName": "짜장면",
*		"userId": 1,
*		"score": 5
*	}
 */
func PostMenuHandler(c *gin.Context) {

	// 요청 바디를 Menu table 뿐만 아니라 User table과도 연관된 Score를 저장할 수 있는 구조체로 파싱
	var request struct {
		MenuName string `json:"menuName"`
		UserID   int64  `json:"userId"`
		Score    int    `json:"score"`
	}

	// 요청 바디를 파싱하여 request 구조체에 저장
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// findMenu 함수를 통해 요청 바디에 담긴 메뉴 이름을 가진 Menu를 찾음
	// 요청 바디에 담긴 메뉴 이름을 가진 Menu가 없으면 insertMenu 함수를 통해 Menu table에 새로운 메뉴를 삽입
	// 요청 바디에 담긴 메뉴 이름을 가진 Menu가 있으면 updateMenu 함수를 통해 Menu table에 메뉴의 평점을 갱신
	_, err := service.FindMenuByName(request.MenuName)
	if err != nil {
		// 값이 없으면 insertMenu
		service.InsertMenu(request.MenuName, request.Score)
	} else {
		service.UpdateMenu(request.MenuName, request.Score)
	}

	// findUser 함수를 통해 요청 바디에 담긴 사용자 ID를 가진 User를 찾음
	// 요청 바디에 담긴 사용자 ID를 가진 User가 없으면 값이 없다는 에러를 반환
	// 요청 바디에 담긴 사용자 ID를 가진 User가 있으면 updateMenuScore 함수를 통해 User table에 메뉴의 평점을 갱신
	user, err := service.FindUserByID(request.UserID)
	if err != nil {
		// 값이 없으면 에러 처리
		fmt.Println("Invalid user id: %v", err)
	} else {
		service.UpdateMenuScore(user, request.MenuName, request.Score)
	}
}
