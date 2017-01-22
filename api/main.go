package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func randomColor() string {
	return fmt.Sprintf("#%.2x%.2x%.2x", rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

func encodeResponse(w http.ResponseWriter, object interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(object); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/api/v1/colors", func(w http.ResponseWriter, r *http.Request) {
		var colors []string
		for i := 0; i < 20; i++ {
			colors = append(colors, randomColor())
		}
		encodeResponse(w, colors)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
