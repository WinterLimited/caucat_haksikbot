package model

import "time"

// Haksik: 학식 정보 구조체
type Haksik struct {
	Date       time.Time `json:"date"`
	Camp       string    `json:"camp"`
	Time       string    `json:"time"`
	Course     string    `json:"course"`
	Menu       []string  `json:"menu"`
	Picture    string    `json:"picture"`
	Price      int       `json:"price"`
	Restaurant string    `json:"restaurant"`
}

// FetchData: fetch 함수의 요청 데이터 구조체
type FetchData struct {
	TimeOfDay int
	Days      int
	IsSeoul   bool
}

// RequestData: request 함수의 요청 데이터 구조체
type RequestData struct {
	Tabs  string `json:"tabs"`
	Tabs2 string `json:"tabs2"`
	Daily int    `json:"daily"`
}

// ResponseData: request 함수의 응답 데이터 구조체
type APIResponse struct {
	IsEmpty string `json:"isEmpty"`
	List    []struct {
		Camp       string `json:"camp"`
		Course     string `json:"course"`
		Date       string `json:"date"`
		MenuDetail string `json:"menuDetail"`
		PicPath    string `json:"picPath"`
		Price      string `json:"price"`
		Rest       string `json:"rest"`
		Time       string `json:"time"`
	} `json:"list"`
}
