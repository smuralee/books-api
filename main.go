package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"smuralee.com/books-api/pkg/apis/books"
	"smuralee.com/books-api/pkg/apis/external"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	books.Handler(router)
	external.Handler(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
