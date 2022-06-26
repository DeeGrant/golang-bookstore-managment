package main

import (
	"fmt"
	"github.com/DeeGrant/golang-bookstore-management/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	fmt.Println("Server listening on port 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
