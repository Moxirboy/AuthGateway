package main

import (
	"Auth_service-and-Api_gateway/proto"
	"Auth_service-and-Api_gateway/service/conn"
	"Auth_service-and-Api_gateway/service/service/repo"
	"Auth_service-and-Api_gateway/service/service/usecase"
	"google.golang.org/grpc"
	"log"
	"net"
)

var port = "localhost:3000"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	con := repo.NewAuthPostgres(conn.Db)
	server := usecase.NewServer(con)
	proto.RegisterAuthServer(s, server)
	log.Panicf("Server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v ", err)
	}
}
