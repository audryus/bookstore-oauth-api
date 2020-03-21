package cassandra

import (
	"os"

	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	cluster := gocql.NewCluster(os.Getenv("cassandra_host"))
	cluster.Keyspace = os.Getenv("cassandra_keyspace")
	cluster.Consistency = gocql.Quorum
	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

//GetSession from cassandra
func GetSession() *gocql.Session {
	return session
}
