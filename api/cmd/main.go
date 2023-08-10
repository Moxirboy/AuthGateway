package main

import "google.golang.org/grpc"

var Address = "localhost:3000"

func main() {
	ClientConn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

}
