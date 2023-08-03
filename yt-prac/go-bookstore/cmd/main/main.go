package main

/*
	docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -e MYSQL_DATABASE=my_db -d -p 3306:3306 mysql:8.0.34-debian
*/

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raphhawk/bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	log.Println("Listening at port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
