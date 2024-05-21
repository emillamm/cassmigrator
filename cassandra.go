package main

import (
	"github.com/gocql/gocql"
)

type Cassandra struct {
	Session *gocql.Session
}

func NewCassandra(
	user string,
	pass string,
	host string,
) (*Cassandra, error) {
	cluster := gocql.NewCluster(host)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: user,
		Password: pass,
	}
	session, err := cluster.CreateSession()
	if err != nil {
		// Currently, gocql doesn't return a typed error when authentication fails.
		// So we create our own based on the error string.
		// See https://github.com/gocql/gocql/blob/v1.6.0/session.go#L187
		err = ParseErrorString(err.Error())
	}
	return &Cassandra{session}, err
}

//func (c *CassandraConnector) NewCassandraConnector CassandraConnector, {
//}

