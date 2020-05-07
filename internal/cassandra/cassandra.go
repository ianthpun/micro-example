package cassandra

import (
	"github.com/gocql/gocql"
)

type Config struct {
	Host     string
	Keyspace string
}

func Connect(c Config) (*gocql.Session, error) {
	cluster := gocql.NewCluster(c.Host)
	cluster.Keyspace = c.Keyspace
	return cluster.CreateSession()
}
