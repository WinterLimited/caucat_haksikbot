package repository

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"server/model"
	"server/store"
	"strconv"
	"strings"
	"sync"
	"time"
)

// timeOfDay: 0(아침), 1(점심), 2(저녁)
// dateFromToday: 오늘로부터 며칠 뒤인지(0이면 오늘, 1이면 내일, ...)
// isSeoul: 서울캠이면 true, 안면캠이면 false
func request(timeOfDay int, dateFromToday int, isSeoul bool) ([]model.Haksik, error) {
	url := "https://mportal.cau.ac.kr/portlet/p005/p005.ajax"
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json;charset=UTF-8",
	}

	tabs := "2"
	if isSeoul {
		tabs = "1"
	}

	data := &model.RequestData{
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

	var responseData model.APIResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, err
	}

	if responseData.IsEmpty == "N" {
		var haksikList []model.Haksik
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

			haksikList = append(haksikList, model.Haksik{
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

/**
 * @api {post} /haksi 학식 정보 요청
 * fetch 함수는 request 함수를 호출하고, 결과를 HakiskInfo에 저장합니다.
 * *sync.WaitGroup: 고루틴이 모두 끝날 때까지 기다리기 위해 사용합니다.
 */
func Fetch(timeOfDay int, days int, isSeoul bool, wg *sync.WaitGroup) {
	defer wg.Done() // 고루틴이 끝나면 wg.Done()을 호출하여 고루틴이 끝났음을 알립니다.

	haksikList, err := request(timeOfDay, days, isSeoul)
	if err != nil {
		log.Fatal(err)
	}
	for _, haksik := range haksikList {
		store.HaksikInfo = append(store.HaksikInfo, haksik)
	}
}
