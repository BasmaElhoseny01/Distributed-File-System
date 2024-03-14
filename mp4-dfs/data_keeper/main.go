package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
	"sync"
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
	fmt.Println("Data Keeper ID: ", id)
	// TODO (2) Send Alive pings to the master node
	// send alive pings to the master node
	alive := Reg.NewHeartBeatServiceClient(connToMaster)
	wg := sync.WaitGroup{}
	// add 1 goroutines to the wait group
	wg.Add(1)
	go func(client Reg.HeartBeatServiceClient, id string) {
		defer wg.Done()
		for {
			_, err := client.AlivePing(context.Background(), &Reg.AlivePingRequest{DataKeeperId: id})
			if err != nil {
				fmt.Println(err)
			}
			// sleep for 1 seconds
			time.Sleep(1 * time.Second)
			fmt.Println("Alive Ping Sent")
		}
	}(alive, id)
	wg.Wait()
}
