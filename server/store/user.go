package store

import (
	"github.com/gocql/gocql"
	"server/model"
)

type UserStore struct {
	Session *gocql.Session
}

func NewUserStore(session *gocql.Session) *UserStore {
	return &UserStore{Session: session}
}

/**
 * @description: store(scylladb)에 저장된 사용자 정보를 반환합니다.
 * @return: []model.Haksik, error
 */
func (s *UserStore) GetUser() ([]model.User, error) {
	iter := s.Session.Query("SELECT * FROM user").Iter()

	var users []model.User
	var user model.User
	for iter.Scan(&user.ID, &user.Name) {
		users = append(users, user)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return users, nil
}

/*
*
* 요청 URL에 담긴 사용자 ID를 통해 사용자의 이름과 메뉴별 평점을 반환
TODO: 여기서부터 작업
*/
func (s *UserStore) GetUserByID(id int64) (*model.User, error) {
	// 사용자 이름 조회
	var user model.User
	if err := s.Session.Query("SELECT * FROM user WHERE id = ?", id).Scan(&user.ID, &user.Name); err != nil {
		return nil, err
	}

	// 메뉴별 평점 조회
	var menuScores []model.MenuScore
	iter := s.Session.Query("SELECT * FROM menu_score WHERE user_id = ?", id).Iter()
	var menuScore model.MenuScore
	for iter.Scan(&menuScore.UserId, &menuScore.MenuName, &menuScore.Score) {
		menuScores = append(menuScores, menuScore)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return &user, nil
}
