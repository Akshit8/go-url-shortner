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
func NewURLRepository(session *gocql.Session, tableName string) urls.Repository {
	return &urlRepository{session: session, tableName: tableName}
}

func (u *urlRepository) Save(url *urls.URL) (*urls.URL, error) {
	return nil, nil
}

func (u *urlRepository) Get(code string) (*urls.URL, error) {
	return nil, nil
}
