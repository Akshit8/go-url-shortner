package cassandra

import (
	"github.com/Akshit8/url-shortner/pkg/redirect"
	"github.com/gocql/gocql"
)

type redirectRepository struct {
	session   *gocql.Session
	tableName string
}

// NewRedirectRepository creates new instance of redirectRepository
func NewRedirectRepository(session *gocql.Session, tableName string) redirect.Repository {
	return &redirectRepository{session: session, tableName: tableName}
}

func (r *redirectRepository) Find(code string) (string, error) {
	return "", nil
}
