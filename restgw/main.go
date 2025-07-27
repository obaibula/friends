package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	gw "github.com/obaibula/friends/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const grpc_server_addr = "localhost:8080"

func main() {
	ctx := context.TODO()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterFriendServiceHandlerFromEndpoint(ctx, mux, grpc_server_addr, opts)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("http listening on 8081")
	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}
