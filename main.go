package main

import (
	"fmt"
	"main/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)

	fmt.Println("Listening on port 8080")
	http.Handle("/", r)

	http.ListenAndServe(":8080", r)
}
