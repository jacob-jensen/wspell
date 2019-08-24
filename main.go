package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jacob-jensen/wspell/word"
)

func handleWord(w http.ResponseWriter, r *http.Request) {
	data := word.Word{Word: "SØSTJERNE", Category: 1}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func handleWordID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("id is:", vars["id"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/word", handleWord).Methods("Get")
	r.HandleFunc("/word/{id}", handleWordID).Methods("Get")
	http.ListenAndServe(":8080", r)
}
