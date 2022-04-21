package main

import (
	"context"
	"log"
	"net"

	pb "github.com/LuizFJP/currency-coin-grpc-go/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CurrencyCoinServiceServer
}

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))

	if err != nil {
		log.Fatalf(err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatalf(err)
	}

	collection = client.Database("coinsdb").Collection("coins")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCurrencyCoinServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to server %v\n", err)
	}

}