package server

import (
	"fmt"
	"net"

	DB "go.videos.core/db"
	"google.golang.org/grpc"
)

func StartServer(addr string) {
	baseServer := grpc.NewServer()
	authDB, _ := DB.InitDB()
	registerAuthServer(baseServer, authDB)
	runServer(baseServer, addr)
}

func runServer(server *grpc.Server, addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Starting server at: %+v\n", addr)
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
