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

package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

var Books = initializeBooks()

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Printf("Fetching all the books...\n")

	err := json.NewEncoder(w).Encode(Books)
	if err != nil {
		panic(err)
	}
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Printf("Fetching a book by id - %s ...\n", id)

	for _, book := range Books {
		if book.Id == id {
			err := json.NewEncoder(w).Encode(book)
			if err != nil {
				panic(err)
			}
		}
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Printf("Adding a book...\n")

	reqBody, _ := io.ReadAll(r.Body)
	var book Book

	marshalErr := json.Unmarshal(reqBody, &book)
	if marshalErr != nil {
		panic(marshalErr)
	}
	Books = append(Books, book)

	encodeErr := json.NewEncoder(w).Encode(book)
	if encodeErr != nil {
		panic(encodeErr)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Printf("Updating a book using the id - %s ...\n", id)

	reqBody, _ := io.ReadAll(r.Body)
	var updatedBook Book

	marshalErr := json.Unmarshal(reqBody, &updatedBook)
	if marshalErr != nil {
		panic(marshalErr)
	}

	for index, book := range Books {
		if book.Id == id {
			fmt.Printf("Match found for a book update, updating...\n")
			Books = append(Books[:index], Books[index+1:]...)
			Books = append(Books, updatedBook)
			encodeErr := json.NewEncoder(w).Encode(updatedBook)
			if encodeErr != nil {
				panic(encodeErr)
			}
			return
		}
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Printf("Deleting a book by id - %s ...\n", id)

	for index, book := range Books {
		if book.Id == id {
			fmt.Println("Match found for a book delete")
			Books = append(Books[:index], Books[index+1:]...)
			return
		}
	}
}
