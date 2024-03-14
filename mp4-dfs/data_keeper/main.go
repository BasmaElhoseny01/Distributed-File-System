package main

import (
	"fmt"
	"google.golang.org/grpc"
	Reg "mp4-dfs/schema/register"
)


func main() {
	fmt.Println("Hello From Data Node 📑")
	// TODO (1) Register to the master node

	serverPort := "localhost:5002"
	connToMaster, err := grpc.Dial(serverPort, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer connToMaster.Close()

	client := Reg.NewRegisterClient(connToMaster)
	// TODO (2) Send Alive pings to the master node
	


}