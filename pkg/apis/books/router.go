/**
   Copyright 2020 Suraj Muraleedharan

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package books

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	v1 "smuralee.com/books-api/pkg/apis/books/v1"
)

var response FixedResponse

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response.Status = "Success"
	response.Message = "Health check is successful"

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}

func Handler(router *mux.Router) {
	router.HandleFunc("/", welcome)
	router.HandleFunc("/books/v1", v1.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/v1/{id}", v1.GetBookById).Methods("GET")
	router.HandleFunc("/books/v1/", v1.CreateBook).Methods("POST")
	router.HandleFunc("/books/v1/{id}", v1.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/v1/{id}", v1.DeleteBook).Methods("DELETE")
}
