package main

import (
	"github.com/gocql/gocql"
)

func NewPasswordSession(
	user string,
	pass string,
	host string,
	keyspace string,
) (session *gocql.Session, err error) {
	cluster := gocql.NewCluster(host)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: user,
		Password: pass,
	}
	if keyspace != "" {
		cluster.Keyspace = keyspace
	}
	session, err = cluster.CreateSession()
	if err != nil {
		// Currently, gocql doesn't return a typed error when authentication fails.
		// So we create our own based on the error string.
		// See https://github.com/gocql/gocql/blob/v1.6.0/session.go#L187
		err = ParseErrorString(err.Error())
	}
	return
}

