package main

import (
	"context"
	"fmt"
	"net"
	"sync"

	data_lookup "mp4-dfs/master_tracker/data_lookup"
	reg "mp4-dfs/schema/register"

	"google.golang.org/grpc"
)


type masterServer struct {
	reg.UnimplementedDataKeeperRegisterServiceServer
	reg.UnimplementedHeartBeatServiceServer
	data_node_lookup_table data_lookup.DataNodeLookUpTable
}


// DataKeepersNodes Registration Services rpc
func (s *masterServer) Register(ctx context.Context, in *reg.DataKeeperRegisterRequest) (*reg.DataKeeperRegisterResponse, error) {
	fmt.Println("Received: ", in.GetIp())
	fmt.Println("Received: ", in.GetPort())


	// Add the data node to the lookup table
	node_id,err := s.data_node_lookup_table.AddDataNode(&data_lookup.DataNode{Id: ""})
	if err == nil {
		fmt.Printf("New Data Node '%s' added Successfully\n", node_id)
	}

	return &reg.DataKeeperRegisterResponse{DataKeeperId: node_id}, nil
}

// HeartBeat Registration Services rpc
func (s *masterServer) AlivePing(ctx context.Context, in *reg.AlivePingRequest) (*reg.AlivePingResponse, error) {
	fmt.Println("Received: ", in.GetDataKeeperId())
	s.data_node_lookup_table.UpdateNodeStatus( in.GetDataKeeperId(),true)
	return nil,nil
}


func handleClient(master *masterServer) {
	fmt.Println("Handle Client")
	// listen to the port
	dataKeeper_listener, err := net.Listen("tcp", "localhost:5001")
	if err != nil {
		fmt.Println(err)
	}
	defer dataKeeper_listener.Close()

	// define our master server and register the service
	s := grpc.NewServer()
	reg.RegisterDataKeeperRegisterServiceServer(s, master)
	if err := s.Serve(dataKeeper_listener); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Handle Client")
}

func handleDataKeeper(master *masterServer) {
	fmt.Println("Handle Data Keeper")
	// listen to the port
	dataKeeper_listener, err := net.Listen("tcp", "localhost:5002")
	if err != nil {
		fmt.Println(err)
	}
	defer dataKeeper_listener.Close()

	// define our master server and register the service
	s := grpc.NewServer()

	// Register in New Node Registration Service
	reg.RegisterDataKeeperRegisterServiceServer(s,master)

	// Register in HeartBeat Service
	reg.RegisterHeartBeatServiceServer(s,master)
	
	if err := s.Serve(dataKeeper_listener); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Handle Data Keeper")
}

func main() {
	// TODO Thread to listen to alive pings from data keepers
	fmt.Println("Hello From Master Node 😎")
	// Create Master Server
	master:=masterServer{data_node_lookup_table:data_lookup.NewDataNodeLookUpTable()}
	// TODO (1) Register to the master node
	wg := sync.WaitGroup{}
	// add 2 goroutines to the wait group
	wg.Add(2)
	go func() {
		defer wg.Done()
		handleClient(&master)
	}()
	go func() {
		defer wg.Done()
		handleDataKeeper(&master)
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
