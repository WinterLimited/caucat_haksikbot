package model

// User represents a user
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// MenuScore represents a menu score by a user
type MenuScore struct {
	UserId   int64  `json:"userId"`
	MenuName string `json:"menuName"`
	Score    int    `json:"score"`
	Count    int    `json:"count"`
}
