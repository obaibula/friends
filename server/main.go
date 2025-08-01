package main

import (
	"context"
	"log"
	"net"

	pb "github.com/obaibula/friends/proto"
	"google.golang.org/grpc"
)

const (
	uri      = "neo4j://localhost:7687"
	user     = "neo4j"
	password = "your_password"
)

func main() {
	ctx := context.TODO()

	g, err := NewGraph()
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = g.Driver.Close(ctx) }()

	err = g.VerifyConnectivity(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connection established")

	srv := grpc.NewServer()
	defer srv.GracefulStop()

	pb.RegisterFriendServiceServer(srv, NewFriendService(g))

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("gRPC server listening on :8080")

	err = srv.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
