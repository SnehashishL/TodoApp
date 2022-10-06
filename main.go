package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/App/routes"
	_ "github.com/lib/pq"
)

func main() {

	r := routes.RouteInit()
	http.Handle("/", r)

	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8000", r))
}
