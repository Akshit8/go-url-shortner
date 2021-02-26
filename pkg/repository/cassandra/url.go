package cassandra

import (
	"github.com/Akshit8/url-shortner/pkg/url"
	"github.com/gocql/gocql"
)

type urlRepository struct {
	session   *gocql.Session
	tableName string
}

// NewURLRepository creates a new instance of urlRepository
func NewURLRepository(session *gocql.Session, tableName string) url.Repository {
	return &urlRepository{session: session, tableName: tableName}
}

func (u *urlRepository) Save(url *url.URL) (*url.URL, error) {
	return nil, nil
}

func (u *urlRepository) Get(code string) (*url.URL, error) {
	return nil, nil
}
