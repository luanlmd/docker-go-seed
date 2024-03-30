package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Success bool
	Message string
}

func root(res http.ResponseWriter, req *http.Request) {
	message := Response{Success: true, Message: "Hello, World!"}
	b, err := json.Marshal(message)

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(b)
}

func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe(":3000", nil)
	fmt.Println("Hello, World!")
}
