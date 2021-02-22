package cassandra

import "github.com/gocql/gocql"

// client struct defines client for mongodb
type client struct{}

// NewClient creates a instance of client and returns cassandra session
func NewClient(cassandraHost string, cassandraPort int, cassandraKeyspace string) (*gocql.Session, error) {
	cluster := gocql.NewCluster(cassandraHost)
	cluster.Port = cassandraPort
	cluster.Keyspace = cassandraKeyspace
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}
