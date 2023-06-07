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

package mock

import (
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

func getHttpStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	codes := [3]int{200, 403, 500}
	index := rand.Intn(len(codes))
	randomCode := codes[index]

	fmt.Printf("Fetching the response for status code - %d ...\n", randomCode)

	resp, err := http.Get("https://httpstat.us/" + strconv.Itoa(randomCode))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Client: status code: %d\n", resp.StatusCode)
	w.WriteHeader(resp.StatusCode)
}

func Handler(router *mux.Router) {
	router.HandleFunc("/mock", getHttpStatus).Methods("GET")
}
