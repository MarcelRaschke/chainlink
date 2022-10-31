package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleReports(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		fmt.Printf("%s is not supported\n", req.Method)
		return
	}

	fmt.Println("POST /reports called")

	decoder := json.NewDecoder(req.Body)

	var body interface{}

	err := decoder.Decode(&body)
	if err != nil {
		fmt.Printf("failed to decode body %v\n", err)
		return
	}

	fmt.Printf("received %v\n", body)
}

func main() {
	http.HandleFunc("/reports", handleReports)

	fmt.Println("running server on :3000")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
