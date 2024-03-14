package main

import (
	"context"
	"fmt"
	data_lookup "mp4-dfs/master_tracker/data_lookup"
	reg "mp4-dfs/schema/register"
	"net"
	"sync"
	"google.golang.org/grpc"
)

// var data_node_lookup_table=NewDataNodeLookUpTable()
// var myVariable MyStruct // Declaring myVariable globally within the package

type masterServer struct {
	reg.UnimplementedDataKeeperRegisterServiceServer
	data_node_lookup_table data_lookup.DataNodeLookUpTable
}

func (s *masterServer) Register(ctx context.Context, in *reg.DataKeeperRegisterRequest) (*reg.DataKeeperRegisterResponse, error) {
	fmt.Println("Received: ", in.GetIp())
	fmt.Println("Received: ", in.GetPort())

	// Add the data node to the lookup table
	node_id, err := s.data_node_lookup_table.AddDataNode(&data_lookup.DataNode{Id: ""})
	if err == nil {
		fmt.Printf("New Data Node '%s' added Successfully\n", node_id)
	}

	return &reg.DataKeeperRegisterResponse{DataKeeperId: node_id}, nil
}

func handleClient() {
	fmt.Println("Handle Client")
	// listen to the port
	dataKeeper_listener, err := net.Listen("tcp", "localhost:5001")
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
	fmt.Println("Handle Client")
}

func handleDataKeeper() {
	fmt.Println("Handle Data Keeper")
	// listen to the port
	dataKeeper_listener, err := net.Listen("tcp", "localhost:5002")
	if err != nil {
		fmt.Println(err)
	}
	defer dataKeeper_listener.Close()

	// define our master server and register the service
	s := grpc.NewServer()

	reg.RegisterDataKeeperRegisterServiceServer(s, &masterServer{
		data_node_lookup_table: data_lookup.NewDataNodeLookUpTable(),
	})

	if err := s.Serve(dataKeeper_listener); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Handle Data Keeper")
}

func main() {
	// TODO Thread to listen to alive pings from data keepers
	fmt.Println("Hello From Master Node 😎")
	// TODO (1) Register to the master node
	wg := sync.WaitGroup{}
	// add 2 goroutines to the wait group
	wg.Add(2)
	go func() {
		defer wg.Done()
		handleClient()
	}()
	go func() {
		defer wg.Done()
		handleDataKeeper()
	}()

	// wait for all goroutines to finish
	wg.Wait()

	// // listen to the port
	// dataKeeper_listener, err := net.Listen("tcp", "localhost:5002")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer dataKeeper_listener.Close()

	// // define our master server and register the service
	// s := grpc.NewServer()
	// reg.RegisterDataKeeperRegisterServiceServer(s, &masterServer{})
	// if err := s.Serve(dataKeeper_listener); err != nil {
	// 	fmt.Println(err)
	// }
}
