package database

import (
	"log"

	"github.com/gocql/gocql"
)

type DBConnection struct {
	session *gocql.Session
}

func SetupDBConnection() *DBConnection {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Consistency = gocql.Quorum
	cluster.Keyspace = "first_app"
	session, _ := cluster.CreateSession()
	return &DBConnection{session: session}
}

func (db *DBConnection) ExecuteQuery(query string, values ...interface{}) error {
	if err := db.session.Query(query).Bind(values...).Exec(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
