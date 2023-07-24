// store/store.go
package store

import "github.com/gocql/gocql"

type Store struct {
	HaksikStore *HaksikStore
	UserStore   *UserStore
}

func NewStore(session *gocql.Session) *Store {
	return &Store{
		HaksikStore: NewHaksikStore(session),
		UserStore:   NewUserStore(session),
	}
}
