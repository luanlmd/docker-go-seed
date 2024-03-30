package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Response struct {
	Success bool
	Message string
}

func root(res http.ResponseWriter, req *http.Request) {
	message := Response{Success: true, Message: "Hello from Golang!"}
	b, err := json.Marshal(message)

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(b)
}

func main() {
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)

	http.HandleFunc("/", root)
	http.ListenAndServe(addr, nil)
}
