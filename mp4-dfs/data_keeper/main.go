package main

import (
	"fmt"
	"context"
	"google.golang.org/grpc"
	Reg "mp4-dfs/schema/register"
)

var id string

func main() {
	fmt.Println("Hello From Data Node 📑")
	// TODO (1) Register to the master node

	serverPort := "localhost:5002"
	connToMaster, err := grpc.Dial(serverPort, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer connToMaster.Close()

	client := Reg.NewDataKeeperRegisterServiceClient(connToMaster)

	// TODO (1) Register to the master node
	response, err := client.Register(context.Background(), &Reg.DataKeeperRegisterRequest{Ip: "localhost", Port: "5003"})
	if err != nil {
		fmt.Println(err)
	}
	// print the response
	fmt.Println("Registered to the master node")
	id = response.GetDataKeeperId()
	// TODO (2) Send Alive pings to the master node	
}