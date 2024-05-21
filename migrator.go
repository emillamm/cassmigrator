package main

type Migrator struct {
	provider MigrationProvider
	//cassandra Cassandra
	AllowUpdatePassword bool
}

//func (m *Migrator) Run() error {
//	//if m.AllowUpdatePassword {
//	//}
//}

