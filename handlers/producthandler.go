package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi.go/models"


	"github.com/gocql/gocql"
)
	
var Session *gocql.Session

func init() {
	var err error
	
	cluster := gocql.NewCluster("127.0.0.1")
    cluster.ProtoVersion = 3
	cluster.Keyspace = "ecommerce"
    Session, err =cluster.CreateSession()
    if err !=nil {
		panic(err)
    }
    fmt.Println("cassandra initialized")
}

func Getallproducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Products
	m :=map[string]interface{}{}
	iter := Session.Query("SELECT * FROM products").Iter()
	for iter.MapScan(m) {
		products = append(products, models.Products{
			Id:          m["id"].(int),
			Name:        m["name"].(string),
			Description: m["description"].(string),
			Category:    m["category"].(string),
			Price:       m["price"].(float64),
		})
		m = map[string]interface{}{}
	}

	Conv, _ := json.MarshalIndent(products, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))
	
}


