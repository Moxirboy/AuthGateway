package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

var port = "localhost:3000"

type server struct {
	proto.AuthServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.registerAuthServer(s, &server{})
	log.Panicf("Server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v ", err)
	}
}
