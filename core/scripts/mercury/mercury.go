package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/smartcontractkit/wsrpc/peer"
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

type pingServer struct{}

func (s *pingServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	// Extracts the connection client's public key.
	// You can use this to identify the client
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, errors.New("could not extract public key")
	}
	pubKey := p.PublicKey

	fmt.Println(pubKey)

	return &pb.PingResponse{
		Body: "Pong",
	}, nil
}
