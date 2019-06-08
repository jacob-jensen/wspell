package main

import (
	"encoding/json"
	"net/http"
)

//Word contains a single word
type Word struct {
	Word        string `json:"word"`
	Category    int    `json:"category"`
	PicturePath string `json:"picturePath"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		words := []Word{
			Word{Word: "hej", Category: 1},
			Word{Word: "med dig", Category: 2},
		}
		if err := json.NewEncoder(w).Encode(words); err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", mux)
}
