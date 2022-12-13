package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"smuralee.com/books-api/pkg/apis/books"
	"smuralee.com/books-api/pkg/apis/mock"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	books.Handler(router)
	mock.Handler(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
