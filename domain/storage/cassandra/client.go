package cassandra

import (
	"os"
	"time"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func Connect() error {
	host := os.Getenv("CASSANDRA_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("CASSANDRA_PORT")
	if port == "" {
		port = "9042"
	}

	cluster := gocql.NewCluster(host)
	cluster.Port = 9042
	cluster.Keyspace = "apartments"
	cluster.Consistency = gocql.Quorum
	cluster.Timeout = 5 * time.Second
	cluster.ConnectTimeout = 5 * time.Second

	var err error
	Session, err = cluster.CreateSession()
	return err
}
