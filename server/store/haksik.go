package store

import (
	"github.com/gocql/gocql"
	"server/model"
)

// HaksikStore는 store.Store 구조체의 HakiskStore 필드에 대한 메서드를 정의합니다.
type HaksikStore struct {
	Session *gocql.Session
}

func NewHaksikStore(session *gocql.Session) *HaksikStore {
	return &HaksikStore{Session: session}
}

/**
 * @description: store(scylladb)에 저장된 학식 정보를 반환합니다.
 * @return: []model.Haksik, error
 */
func (s *HaksikStore) GetHaksik() ([]model.Haksik, error) {
	iter := s.Session.Query("SELECT * FROM haksik").Iter()

	var haksiks []model.Haksik
	var haksik model.Haksik
	for iter.Scan(&haksik.Camp, &haksik.Course, &haksik.Date, &haksik.Menu, &haksik.Picture, &haksik.Price, &haksik.Restaurant, &haksik.Time) {
		haksiks = append(haksiks, haksik)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return haksiks, nil
}
