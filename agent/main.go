package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jufianto/state-agent/agent/client"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	srvGrpc := grpc.NewServer()

	srv := client.NewAgentClient("key123")
	srv.RegisterGW(srvGrpc)

	fmt.Println("service listening....")

	net, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatalf("failed to register grpc: %v", err)
	}

	reflection.Register(srvGrpc)
	srvGrpc.Serve(net)

}
