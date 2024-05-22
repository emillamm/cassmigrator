package main

import (
	"github.com/gocql/gocql"
	"fmt"
)

//type Migrator struct {
//	provider MigrationProvider
//	//cassandra Cassandra
//	AllowUpdatePassword bool
//}

//func (m *Migrator) Run() error {
//	//if m.AllowUpdatePassword {
//	//}
//}

func RunMigrations(
	session *gocql.Session,
	migrations []Migration,
) error {
	query := fmt.Sprintf("create table if not exists migrations(id text, ts timestamp, primary key (id))")
	if err := session.Query(query).Exec(); err != nil {
		return err
	}
	return nil
}

