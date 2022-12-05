package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandlerFunc)
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		v := map[string]any{
			"hello": "world",
		}
		json.NewEncoder(w).Encode(v)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func helloHandlerFunc(rw http.ResponseWriter, req *http.Request) {
	res := map[string]any{
		"say": "Hello",
		"to":  "Gophers",
	}
	json.NewEncoder(rw).Encode(res)
}
