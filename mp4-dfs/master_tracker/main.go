package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	reg "mp4-dfs/schema/register"
	"net"
)

type masterServer struct {
	reg.UnimplementedDataKeeperRegisterServiceServer
}

func (s *masterServer) Register(ctx context.Context, in *reg.DataKeeperRegisterRequest) (*reg.DataKeeperRegisterResponse, error) {
	fmt.Println("Received: ", in.GetIp())
	fmt.Println("Received: ", in.GetPort())
	return &reg.DataKeeperRegisterResponse{DataKeeperId: "1234"}, nil
}

func main() {
	// TODO Thread to listen to alive pings from data keepers
	fmt.Println("Hello From Master Node 😎")
	// TODO (1) Register to the master node

	dataKeeper_listener, err := net.Listen("tcp", "localhost:5002")
	if err != nil {
		fmt.Println(err)
	}
	defer dataKeeper_listener.Close()

	// define our master server and register the service
	s := grpc.NewServer()
	reg.RegisterDataKeeperRegisterServiceServer(s, &masterServer{})
	if err := s.Serve(dataKeeper_listener); err != nil {
		fmt.Println(err)
	}
}
