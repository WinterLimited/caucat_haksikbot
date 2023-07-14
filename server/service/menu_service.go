package service

import (
	"fmt"
	"server/model"
	"server/store"
)

// Menu table 모두 반환
func FindMenus() ([]model.Menu, error) {
	return store.Menus, nil
}

// Menu table에 값 존재여부 확인
// error handling 추가
func FindMenuByName(menuName string) (*model.Menu, error) {
	for _, menu := range store.Menus {
		if menu.Name == menuName {
			return &menu, nil
		}
	}
	return nil, fmt.Errorf("Menu not found")
}

// Menu table에 새로운 메뉴 삽입
func InsertMenu(menuName string, score int) (*model.Menu, error) {
	menu := model.Menu{
		Name:  menuName,
		Score: score,
		Count: 1,
	}
	store.Menus = append(store.Menus, menu)
	return &menu, nil
}

// Menu table에 메뉴의 평점 갱신
// *Menu로 값을 직접 받아서 갱신
func UpdateMenu(menuName string, score int) (*model.Menu, error) {
	// 해당 값을 찾아서 갱신
	for i := range store.Menus {
		if store.Menus[i].Name == menuName {
			store.Menus[i].Score = (store.Menus[i].Score*store.Menus[i].Count + score) / (store.Menus[i].Count + 1)
			store.Menus[i].Count++
			return &store.Menus[i], nil
		}
	}
	return nil, fmt.Errorf("Menu not found")
}
