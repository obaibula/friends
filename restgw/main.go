package main

import (
	"context"
	"log"
	"net/http"
	"time"

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

	srv := &http.Server{
		Addr:         ":8081",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Println("http listening on 8081")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
