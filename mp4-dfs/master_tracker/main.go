package main

import (
	"context"
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"

	data_lookup "mp4-dfs/master_tracker/data_lookup"

	req "mp4-dfs/schema/file_transfer_request"
	hb "mp4-dfs/schema/heart_beat"
	reg "mp4-dfs/schema/register"
)


type masterServer struct {
	reg.UnimplementedDataKeeperRegisterServiceServer
	hb.UnimplementedHeartBeatServiceServer
	req.UnimplementedFileTransferRequestServiceServer

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
func (s *masterServer) AlivePing(ctx context.Context, in *hb.AlivePingRequest) (*hb.AlivePingResponse, error) {
	node_id:=in.GetDataKeeperId()
	stamp,err:=s.data_node_lookup_table.UpdateNodeTimeStamp(node_id)

	if err == nil {
		fmt.Printf("Data Node '%s' Ping Time stamp Updated with %s \n", node_id,stamp.Format("2006-01-02 15:04:05"))
	}
	return &hb.AlivePingResponse{},nil
}

// FileTransferRequest Services rpc
func (s *masterServer) UploadRequest (ctx context.Context, in *req.UploadFileRequest) (*req.UploadFileResponse,error){
	fmt.Println("Received Upload Request")
	return  &req.UploadFileResponse{Host: "localhost",Port: "5003"},nil
}


func handleClient(master *masterServer) {
	fmt.Println("Handle Client")
	// listen to the port
	client_listener, err := net.Listen("tcp", "localhost:5001")
	if err != nil {
		fmt.Println(err)
	}
	defer client_listener.Close()

	// define our master server and register the service
	s := grpc.NewServer()

	// Register in File Transfer Request Service
	req.RegisterFileTransferRequestServiceServer(s, master)

	if err := s.Serve(client_listener); err != nil {
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
	hb.RegisterHeartBeatServiceServer(s,master)
	
	if err := s.Serve(dataKeeper_listener); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Handle Data Keeper finished")
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
