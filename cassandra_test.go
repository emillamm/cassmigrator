package main

import (
	"testing"
	"os"
	"github.com/gocql/gocql"
	"slices"
)

func TestCassandra(t *testing.T) {

	host := os.Getenv("CASSANDRA_HOST")
	if host == "" {
		host = "localhost:9042"
	}

	t.Run("NewPasswordSession returns authentication error if incorrect password is used", func(t *testing.T) {
		session, err := NewPasswordSession("cassandra", "invalid_password", host, "")
		defer closeSession(session)
		if _, ok := err.(*InvalidCredentialsError); ok == false {
			t.Errorf("%#v was not the expected error. host %s", err, host)
		}
	})

	t.Run("NewPasswordSession returns a working session with correct creds", func(t *testing.T) {
		session, err := NewPasswordSession("cassandra", "cassandra", host, "")
		defer closeSession(session)
		if err != nil {
			t.Errorf("session creation error not nil: %s", err)
		}
		err = session.Query("describe keyspaces").Exec()
		if err != nil {
			t.Errorf("query error not nil: %s", err)
		}
	})

	t.Run("NewPasswordSession returns a session not tied to a keyspace if keyspace is blank", func(t *testing.T) {
		session, _ := NewPasswordSession("cassandra", "cassandra", host, "")
		defer closeSession(session)
		keyspaces := getKeyspaces(t, session)
		for _, keyspaceToCheck := range []string{"system", "system_auth"} {
			if !slices.Contains(keyspaces, keyspaceToCheck) {
				t.Errorf("expected keyspace %s was not returned by query", keyspaceToCheck)
				break
			}
		}
	})

	t.Run("NewPasswordSession returns a working session tied to a keyspace if keyspace is non-blank", func(t *testing.T) {
		session, _ := NewPasswordSession("cassandra", "cassandra", host, "system")
		defer closeSession(session)
		keyspaces := getKeyspaces(t, session)
		if len(keyspaces) != 1 || keyspaces[0] != "system" {
			t.Errorf("tables returned from unexpected keyspaces %v", keyspaces)
		}
	})
}

func closeSession(session *gocql.Session) {
	if session != nil {
		defer session.Close()
	}
}

func getKeyspaces(t testing.TB, session *gocql.Session) []string {
	t.Helper()
	scanner := session.Query("describe tables").Iter().Scanner()
	var keyspaces []string // keeps track of all keyspaces

	var keyspace_name string
	var ktype string
	var name string
	for scanner.Next() {
		err := scanner.Scan(&keyspace_name, &ktype, &name)
		if err != nil {
			t.Errorf("scan error: %s", err)
		}
		keyspaces = append(keyspaces, keyspace_name)
	}

	if err := scanner.Err(); err != nil {
		t.Errorf("scanner close error: %s", err)
	}

	return slices.Compact(keyspaces)
}

