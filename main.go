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
