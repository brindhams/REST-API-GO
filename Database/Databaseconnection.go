package database

import (
    "log"
    "fmt"

	"github.com/gocql/gocql"
)


type DBConnection struct {
    cluster *gocql.ClusterConfig
    session *gocql.Session
}
var connection DBConnection

func SetupDBConnection() {
    var err error

    connection.cluster = gocql.NewCluster("127.0.0.1")
    connection.cluster.ProtoVersion = 3
    connection.cluster.Keyspace = "ecommerce"
    connection.session, err =connection.cluster.CreateSession()
    if err !=nil {
        log.Fatal(err)
    }
    fmt.Println("cassandra initialized")
    
    
}

