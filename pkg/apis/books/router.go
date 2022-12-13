package books

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	v1 "smuralee.com/books-api/pkg/apis/books/v1"
)

var response FixedResponse

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response.Status = "Success"
	response.RemoteAddress = r.RemoteAddr
	response.Hostname, _ = os.Hostname()

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}

func CatalogHandler() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", welcome)
	myRouter.HandleFunc("/books/v1", v1.GetAllBooks).Methods("GET")
	myRouter.HandleFunc("/books/v1/{id}", v1.GetBookById).Methods("GET")
	myRouter.HandleFunc("/books/v1/", v1.CreateBook).Methods("POST")
	myRouter.HandleFunc("/books/v1/{id}", v1.UpdateBook).Methods("PUT")
	myRouter.HandleFunc("/books/v1/{id}", v1.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
