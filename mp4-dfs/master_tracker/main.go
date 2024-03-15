package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"

	data_lookup "mp4-dfs/master_tracker/data_lookup"
	file_lookup "mp4-dfs/master_tracker/file_lookup"

	cf "mp4-dfs/schema/confirm_file_transfer"
	req "mp4-dfs/schema/file_transfer_request"
	fi "mp4-dfs/schema/finish_file_transfer"
	hb "mp4-dfs/schema/heart_beat"
	reg "mp4-dfs/schema/register"
	download "mp4-dfs/schema/download"
)


type masterServer struct {
	reg.UnimplementedDataKeeperRegisterServiceServer
	hb.UnimplementedHeartBeatServiceServer
	req.UnimplementedFileTransferRequestServiceServer
	fi.UnimplementedFinishFileTransferServiceServer
	cf.UnimplementedConfirmFileTransferServiceServer
	download.UnimplementedDownloadServiceServer

	data_node_lookup_table data_lookup.DataNodeLookUpTable
	files_lookup_table file_lookup.FileLookUpTable
}
func NewMasterServer() masterServer{
	return masterServer{
		data_node_lookup_table:data_lookup.NewDataNodeLookUpTable(),
		files_lookup_table:file_lookup.NewFileLookUpTable(),
	}
}


// DataKeepersNodes Registration Services rpc
func (s *masterServer) Register(ctx context.Context, in *reg.DataKeeperRegisterRequest) (*reg.DataKeeperRegisterResponse, error) {
	fmt.Println("Received: ", in.GetIp())
	fmt.Println("Received: ", in.GetPort())

	// Add the data node to the lookup table
	new_data_node:=data_lookup.NewDataNode(in.GetIp(),in.GetPort())
	node_id, err := s.data_node_lookup_table.AddDataNode(&new_data_node)
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
	file_name:=in.GetFileName()
	fmt.Println("Received Upload Request",file_name)
	// check if file already exist	
	if  exists :=s.files_lookup_table.CheckFileExists(file_name); exists {
		return  &req.UploadFileResponse{},errors.New("file already exists")
    }

	// Get the data node with the least load
	node_address,err:=s.data_node_lookup_table.GetLeastLoadedNode()
	if err != nil {
		fmt.Println(err)
	}
	return  &req.UploadFileResponse{Address: node_address},nil
}

// Finish File Transfer Service rpc
func (s *masterServer) FinishFileUpload(ctx context.Context, in *fi.FinishFileUploadRequest) (*fi.FinishFileUploadResponse, error) {
	fmt.Println("FinishFileUpload  Request")
	data_node_id:=in.GetDataNodeId()
	file_name:=in.GetFileName()
	file_path:=in.GetFilePath()

	// Add File to Files LookUpTable
	newFile:=file_lookup.NewFile(file_name,data_node_id,file_path)
	err:=s.files_lookup_table.AddFile(&newFile)
	if err!=nil{
		fmt.Printf("Error When adding file %s to lookup Table\n",file_name)
		println(err)
		return  &fi.FinishFileUploadResponse{},err
	}
	fmt.Printf("Successfully added File %s at node %s in %s to lookup Table",file_name,data_node_id,file_path)


	// [TODO] Send Notification Message To Client

	// [TODO] Replicas

	return  &fi.FinishFileUploadResponse{},nil
}

// Confirm File Transfer Services rpc
func (s *masterServer) ConfirmFileTransfer (ctx context.Context, in *cf.ConfirmFileTransferRequest) (*cf.ConfirmFileTransferResponse, error){
	file_name:=in.GetFileName();
	 // Try checking the condition 5 times with a 2-second interval
	 for i := 0; i < 5; i++ {
        if exists := s.files_lookup_table.CheckFileExists(file_name); exists {
            return &cf.ConfirmFileTransferResponse{}, nil // File exists, return without error
        }
        time.Sleep(2 * time.Second) // Wait for 2 seconds before checking again
    }
	// If the file doesn't exist after 5 attempts, return an error
	return &cf.ConfirmFileTransferResponse{}, errors.New("file not found")
}

func (s *masterServer) GetServer(ctx context.Context, in *download.DownloadRequest) (*download.DownloadServerResponse, error) {
	file_name:=in.GetFileName()
	fmt.Println("Received Download Request",file_name)
	// check if file already exist
	exists,node_1,_,_,_,_,_ :=s.files_lookup_table.GetFile(file_name)
	if exists == false {
		// return error to client in response data
		return  &download.DownloadServerResponse{Data: &download.DownloadServerResponse_Error{
			Error: "File Not Found",
		}},nil
	}

	Ip,Port:=s.data_node_lookup_table.GetNodeAddress(node_1)
	// return data to client in response data
	// create list of servers which contains ip and port
	servers:= &download.ServerList{
		Servers:[]* download.Server{
			{
				Ip:   Ip,
				Port: Port,
			},
		},
	}
	return  &download.DownloadServerResponse{
		Data: &download.DownloadServerResponse_Servers{
			Servers: servers,
		},
	},nil
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

	download.RegisterDownloadServiceServer(s,master)

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

	// Register to Finish File Transfer Service
	fi.RegisterFinishFileTransferServiceServer(s,master)

	// Register to Confirm File Transfer Service
	cf.RegisterConfirmFileTransferServiceServer(s,master)
	
	download.RegisterDownloadServiceServer(s,master)

	if err := s.Serve(dataKeeper_listener); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Handle Data Keeper finished")
}

func main() {
	// TODO Thread to listen to alive pings from data keepers
	fmt.Println("Hello From Master Node 😎")
	// Create Master Server
	master:=NewMasterServer()
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
}
