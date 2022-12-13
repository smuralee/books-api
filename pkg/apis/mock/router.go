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

	resp, err := http.Get("https://fetch.mock.status/" + strconv.Itoa(randomCode))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Client: status code: %d\n", resp.StatusCode)
	w.WriteHeader(resp.StatusCode)
}

func Handler(router *mux.Router) {
	router.HandleFunc("/mock", getHttpStatus).Methods("GET")
}
