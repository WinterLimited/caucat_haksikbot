package model

// Menu represents a menu item
type Menu struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
	Count int    `json:"count"`
}
