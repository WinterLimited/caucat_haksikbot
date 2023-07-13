package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Haksik: 학식 정보 구조체
type Haksik struct {
	Camp       string
	Course     string
	Date       time.Time
	Menu       []string
	Picture    string
	Price      int
	Restaurant string
	Time       string
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

// HaksikInfo 전달받은 학식 정보를 담는 임시 데이터베이스
var HaksikInfo []Haksik = []Haksik{}

// timeOfDay: 0(아침), 1(점심), 2(저녁)
// dateFromToday: 오늘로부터 며칠 뒤인지(0이면 오늘, 1이면 내일, ...)
// isSeoul: 서울캠이면 true, 안면캠이면 false
func request(timeOfDay int, dateFromToday int, isSeoul bool) ([]Haksik, error) {
	url := "https://mportal.cau.ac.kr/portlet/p005/p005.ajax"
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json;charset=UTF-8",
	}

	tabs := "2"
	if isSeoul {
		tabs = "1"
	}

	data := &RequestData{
		Tabs:  tabs,
		Tabs2: strconv.Itoa(1 << timeOfDay * 10),
		Daily: dateFromToday,
	}

	dataJson, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(dataJson))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// TLS configuration
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			MaxVersion: tls.VersionTLS13,
		},
	}

	client := &http.Client{Transport: transport}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseData APIResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, err
	}

	if responseData.IsEmpty == "N" {
		var haksikList []Haksik
		for _, item := range responseData.List {
			date, err := time.Parse("2006.01.02", item.Date)
			if err != nil {
				return nil, err
			}

			// 가격 정보를 숫자로 변환
			// 예시: "5,000원" -> 5000
			// ,와 원을 제거한 후, 숫자로 변환
			// 문자열의 공백을
			price, err := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(item.Price, ",", ""), "원", ""), " ", ""))
			if err != nil {
				return nil, err
			}

			haksikList = append(haksikList, Haksik{
				Camp:       item.Camp,
				Course:     item.Course,
				Date:       date,
				Menu:       strings.Split(item.MenuDetail, ","),
				Picture:    item.PicPath,
				Price:      price,
				Restaurant: item.Rest,
				Time:       item.Time,
			})
		}

		return haksikList, nil
	}

	return nil, nil
}

// fetch 함수는 request 함수를 호출하고, 결과를 HakiskInfo에 저장합니다.
// *sync.WaitGroup: 고루틴이 모두 끝날 때까지 기다리기 위해 사용합니다.
func fetch(timeOfDay int, days int, isSeoul bool, wg *sync.WaitGroup) {
	defer wg.Done() // 고루틴이 끝나면 wg.Done()을 호출하여 고루틴이 끝났음을 알립니다.

	haksikList, err := request(timeOfDay, days, isSeoul)
	if err != nil {
		log.Fatal(err)
	}
	for _, haksik := range haksikList {
		HaksikInfo = append(HaksikInfo, haksik)
	}
}

// 이하 코드는 각 fetch 함수 및 test 함수를 작성하며 필요에 따라 비슷한 방식으로 작성하면 됩니다.

// main 함수는 테스트용으로 작성되었습니다.
// 실제로 사용할 때는 main 함수를 지우고, fetch 함수를 호출하면 됩니다.
func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	go fetch(0, 0, true, &wg)
	go fetch(1, 0, true, &wg)
	go fetch(2, 0, true, &wg)
	wg.Wait()

	fmt.Println(HaksikInfo)
}
