package main

import (
	"github.com/gocql/gocql"
	"testing"
	"os"
	"fmt"
	"math/rand"
)

func TestMigrate(t *testing.T) {

	//cqlFilesDir := "testdata"
	//migrationProvider := &MigrationProvider{cqlFilesDir}
	user := "cassandra"
	pass := "cassandra"
	host := os.Getenv("CASSANDRA_HOST")
	if host == "" {
		host = "localhost:9042"
	}

	parentSession, _ := NewPasswordSession(user, pass, host, "")
	defer parentSession.Close()

	//t.Run("Run should update password when AllowUpdatePassword is set and authenticating for the first time", func(t *testing.T) {
	//	//m := NewMigrator()
	//	//m.Run()
	//})

	t.Run("RunMigrations should create a migration table if it doesn't exist", func(t *testing.T) {
		ephemeralSession(t, parentSession, user, pass, host, func(session *gocql.Session, keyspace string) {
			// Prepare query
			query := fmt.Sprintf("select table_name from system_schema.tables where keyspace_name='%s' and table_name='migrations'", keyspace)
			checkTableExistence := func(shouldExist bool) {
				iter := session.Query(query).Iter()
				if (iter.NumRows() != 1) && shouldExist {
					t.Errorf("migration table existence=%t in keyspace %s not valid", shouldExist, keyspace)
				}
				iter.Close()
			}

			// Check that table doesn't exist
			checkTableExistence(false)

			// Perform some migration
			migrations := []Migration{
				Migration{
					Id: "001",
					Statements: []string{
						"create table test_table(id text primary key (id))",
					},
				},
			}
			if err := RunMigrations(session, migrations); err != nil {
				t.Errorf("unable to run migrations: %s", err)
			}

			// Check that table does exist
			checkTableExistence(true)
		})
	})

	//t.Run("RunMigrations should execute statements in alphabetical order of id", func(t *testing.T) {
}

func ephemeralSession(
	t testing.TB,
	parentSession *gocql.Session,
	user string,
	pass string,
	host string,
	block func(session *gocql.Session, keyspace string),
) {
	t.Helper()
	var err error
	var session *gocql.Session

	keyspace := randomKeyspace()

	q := fmt.Sprintf("CREATE KEYSPACE %s WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };", keyspace)
	if err = parentSession.Query(q).Exec(); err != nil {
		t.Errorf("failed to create ephemeral keyspace: %s", err)
		return
	}

	defer func() {
		q = fmt.Sprintf("DROP KEYSPACE %s;", keyspace)
		err = parentSession.Query(q).Exec()
		if err != nil {
			t.Errorf("failed to delete ephemeral keyspace %s: %s", keyspace, err)
		}
		session.Close()
	}()

	if session, err = NewPasswordSession(user, pass, host, keyspace); err != nil {
		t.Errorf("failed to create ephemeral session: %s", err)
		return
	}

	block(session, keyspace)
}

// Generates keyspace name in the form of "test_[a-z]7" e.g. test_hqbrluz
func randomKeyspace() string {
	chars := "abcdefghijklmnopqrstuvwxyz"
	length := 7
	b := make([]byte, length)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return fmt.Sprintf("test_%s", string(b))
}

//func newEphemeralSession(
//	t testing.TB,
//	parentSession *gocql.Session,
//	keyspace string,
//	user string,
//	pass string,
//	host string,
//) (session *gocql.Session, closeHandle func()) {
//	var err error
//	q := fmt.Sprintf("CREATE KEYSPACE %s WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };", keyspace)
//	if err = parentSession.Query(q).Exec(); err != nil {
//		t.Errorf("failed to create ephemeral keyspace: %s", err)
//		return
//	}
//	if session, err = NewPasswordSession(user, pass, host, keyspace); err != nil {
//		t.Errorf("failed to create ephemeral session: %s", err)
//		return
//	}
//	closeHandle = func() {
//		session.Close()
//		q = fmt.Sprintf("DROP KEYSPACE %s;", keyspace)
//		err = parentSession.Query(q).Exec()
//		if err != nil {
//			t.Errorf("failed to delete keyspace %s: %s", keyspace, err)
//			return
//		}
//	}
//	return
//}


