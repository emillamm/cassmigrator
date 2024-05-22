package main

import (
	"testing"
	"reflect"
)

func TestMigrationProvider(t *testing.T) {

	provider := &FileMigrationProvider{"testdata"}

	t.Run("read migrations from files", func(t *testing.T) {
		got := provider.GetMigrations()
		want := []Migration{
			Migration{
				Id: "000",
				Statements: []string{
					"CREATE KEYSPACE test_keyspace_1 WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };",
					"CREATE ROLE IF NOT EXISTS test_role_1 WITH PASSWORD = 'test1' AND LOGIN = true;",
					"GRANT ALL ON KEYSPACE test_keyspace_1 TO test_role_1;",
				},
			},
			Migration{
				Id: "001",
				Statements: []string{
					"CREATE KEYSPACE test_keyspace_2 WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };",
					"CREATE ROLE IF NOT EXISTS test_role_2 WITH PASSWORD = 'test2' AND LOGIN = true;",
					"GRANT ALL ON KEYSPACE test_keyspace_2 TO test_role_2;",
				},
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

