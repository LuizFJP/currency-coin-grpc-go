package main

import (
	"log"
	"net"

	pb "github.com/LuizFJP/currency-coin-grpc-go/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

// type Server struct {
	
// }

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
}