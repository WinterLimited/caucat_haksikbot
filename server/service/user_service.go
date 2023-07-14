package service

import (
	"fmt"
	"server/model"
	"server/store"
)

// User table에 값 존재여부 확인
// error handling 추가
func FindUserByID(userID int64) (*model.User, error) {
	for _, user := range store.Users {
		if user.ID == userID {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("User not found")
}

// User table에 메뉴의 평점 갱신
func UpdateMenuScore(user *model.User, name string, score int) {
	// 해당 값을 찾아서 갱신
	for i := range user.MenuScores {
		if user.MenuScores[i].MenuName == name {
			user.MenuScores[i].Score = (user.MenuScores[i].Score*user.MenuScores[i].Count + score) / (user.MenuScores[i].Count + 1)
			user.MenuScores[i].Count++
			return
		}
	}
}
