package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"

	utils "mp4-dfs/utils"

	client_lookup "mp4-dfs/master_tracker/client_lookup"
	data_lookup "mp4-dfs/master_tracker/data_lookup"
	file_lookup "mp4-dfs/master_tracker/file_lookup"

	download "mp4-dfs/schema/download"
	hb "mp4-dfs/schema/heart_beat"
	reg "mp4-dfs/schema/register"
	upload "mp4-dfs/schema/upload"

	cf "mp4-dfs/schema/confirm_file_transfer"
)


type masterServer struct {
	reg.UnimplementedDataKeeperRegisterServiceServer
	hb.UnimplementedHeartBeatServiceServer
	upload.UnimplementedUploadServiceServer
	download.UnimplementedDownloadServiceServer
	
	cf.UnimplementedConfirmFileTransferServiceServer

	data_node_lookup_table data_lookup.DataNodeLookUpTable
	files_lookup_table file_lookup.FileLookUpTable
	client_lookup_table  client_lookup.ClientLookUpTable
}
func NewMasterServer() masterServer{
	return masterServer{
		data_node_lookup_table:data_lookup.NewDataNodeLookUpTable(),
		files_lookup_table:file_lookup.NewFileLookUpTable(),
		client_lookup_table:client_lookup.NewClientLookUpTable(),
	}
}


// DataKeepersNodes Registration Services rpc
func (s *masterServer) Register(ctx context.Context, in *reg.DataKeeperRegisterRequest) (*reg.DataKeeperRegisterResponse, error) {
	// Add the data node to the lookup table
	new_data_node:=data_lookup.NewDataNode(in.GetIp(),in.GetPorts())
	node_id, err := s.data_node_lookup_table.AddDataNode(&new_data_node)
	if err!=nil{
		fmt.Printf("Error When adding new Data Node with ID: %s",node_id)
		return  &reg.DataKeeperRegisterResponse{},err
	}
	fmt.Printf("New Data Node added Successfully\n")
	fmt.Println(s.data_node_lookup_table.PrintDataNodeInfo(node_id))

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

// UploadService RpcS
// RequestUpload rpc
func (s *masterServer) RequestUpload (ctx context.Context, in *upload.RequestUploadRequest) (*upload.RequestUploadResponse,error){
	file_name:=in.GetFileName()

	// check if file already exist	
	if  exists :=s.files_lookup_table.CheckFileExists(file_name); exists {
		fmt.Printf("file %s already exists\n",file_name)
		return  &upload.RequestUploadResponse{},errors.New("")
    }

	// Get the data node with the least load
	node_socket,err:=s.data_node_lookup_table.GetLeastLoadedNode()
	if err != nil {
		fmt.Printf("Can not Get DataNode Port %v\n",err)
		return  &upload.RequestUploadResponse{},err
	}

	// Save The Socket for that client
	client_socket:=in.GetClientSocket()
	err=s.client_lookup_table.AddClient(file_name,client_socket)

	if err!=nil{
		fmt.Printf("Error When adding new Client at %s Waiting for File %s",client_socket,file_name)
		return   &upload.RequestUploadResponse{},err
	}
	fmt.Printf("New Client  added Successfully\n")
	fmt.Println(s.client_lookup_table.PrintClientInfo(file_name))

	return  &upload.RequestUploadResponse{NodeSocket: node_socket},nil
}

// NotifyMaster rpc 
func (s *masterServer) NotifyMaster (ctx context.Context, in *upload.NotifyMasterRequest) (*upload.NotifyMasterResponse,error){
	dataNodeId:=in.GetNodeId()
	fileName:=in.GetFileName()
	filePath:=in.GetFilePath()

	// Add File to Files LookUpTable
	newFile:=file_lookup.NewFile(fileName,dataNodeId,filePath)
	err:=s.files_lookup_table.AddFile(&newFile)
	if err!=nil{
		fmt.Printf("Error When adding file %s to lookup Table\n error %v\n",fileName,err)
		return  &upload.NotifyMasterResponse{},err
	}
	fmt.Printf("New File added Successfully\n")
	fmt.Println(s.files_lookup_table.PrintFileInfo(fileName))

	return &upload.NotifyMasterResponse{}, nil
}

func (s *masterServer) GetServer(ctx context.Context, in *download.DownloadRequest) (*download.DownloadServerResponse, error) {
	file_name:=in.GetFileName()
	fmt.Println("Received Download Request",file_name)
	// check if file already exist
	exists,node_1,path_1,node_2,path_2,node_3,path_3 :=s.files_lookup_table.GetFile(file_name)
	if !exists {
		// return error to client in response data
		return  &download.DownloadServerResponse{Data: &download.DownloadServerResponse_Error{
			Error: "File Not Found",
		}},nil
	}
	Ip1 := ""
	Ip2 := ""
	Ip3 := ""
	Port1 := ""
	Port2 := ""
	Port3 := ""
	// Get the data node with the least load
	if node_1!= "-1"{
		Ip1,Port1 =s.data_node_lookup_table.GetNodeAddress(node_1)
	}
	if node_2!= "-1"{
		Ip2,Port2=s.data_node_lookup_table.GetNodeAddress(node_2)
	}
	if node_3!= "-1"{
		Ip3,Port3 =s.data_node_lookup_table.GetNodeAddress(node_3)
	}
	// create list of servers which contains ip and port
	servers := []*download.Server{
	}
	if Ip1!=""{
		servers = append(servers, &download.Server{
			Ip:   Ip1,
			Port: Port1,
			FilePath: path_1,
		})
	}
	if Ip2!=""{
		servers = append(servers, &download.Server{
			Ip:   Ip2,
			Port: Port2,
			FilePath: path_2,
		})
	}
	if Ip3!=""{
		servers = append(servers, &download.Server{
			Ip:   Ip3,
			Port: Port3,
			FilePath: path_3,
		})
	}
	// create list of servers
	servers_list:= &download.ServerList{
		Servers: servers,
	}
	return  &download.DownloadServerResponse{
		Data: &download.DownloadServerResponse_Servers{
			Servers: servers_list,
		},
	},nil
}

func handleClient(master *masterServer) {
	fmt.Println("Handle Client ")
	// listen to the port
	masterAddress:=utils.GetMasterIP("client")
	client_listener, err := net.Listen("tcp", masterAddress)
	if err != nil {
		fmt.Println(err)
	}
	defer client_listener.Close()
	fmt.Printf("Listening to Client at Socket: %s\n",masterAddress)

	// define our master server and register the service
	s := grpc.NewServer()

	// Register to UploadService
	upload.RegisterUploadServiceServer(s,master)

	// Register to DownloadService
	download.RegisterDownloadServiceServer(s,master)

	if err := s.Serve(client_listener); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Handle Client")
}

func handleDataKeeper(master *masterServer) {
	// listen to the port
	masterAddress:=utils.GetMasterIP("node")
	dataKeeper_listener, err := net.Listen("tcp", masterAddress)
	if err != nil {
		fmt.Println(err)
	}
	defer dataKeeper_listener.Close()
	fmt.Printf("Listening to Data Keeper at Socket: %s\n",masterAddress)

	// define our master server and register the service
	s := grpc.NewServer()

	// Register in New Node Registration Service
	reg.RegisterDataKeeperRegisterServiceServer(s,master)

	// Register to UploadService
	upload.RegisterUploadServiceServer(s,master)

	// Register to DownloadService
	download.RegisterDownloadServiceServer(s,master)

	// Register in HeartBeat Service
	hb.RegisterHeartBeatServiceServer(s,master)
	
	if err := s.Serve(dataKeeper_listener); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Handle Data Keeper finished")
}
func periodicCheckup(master *masterServer){
	for{
		println("Periodic Checkup")
		//1. Check Ideal DataNodes

		// 2.Sent Notifications to Clients
		println("Checking UnConfirmed Files...")
		unconfirmedFiles:=master.files_lookup_table.CheckUnConfirmedFiles()
		println(unconfirmedFiles)
		for _, file := range unconfirmedFiles{
			master.sendClientConfirm(file)

		}
			
			
		// 	//3. Check For Replicas

		// Sleep for 5 seconds before the next check
		time.Sleep(5 * time.Second)
	}

}

func (s *masterServer) sendClientConfirm(fileName string){
	// Send Notification to Client
	// GetSocket for the Client 
	client_socket:=s.client_lookup_table.GetClientSocket(fileName)
	//1. Establish Connection to the Master
	connToClient, err := grpc.Dial(client_socket, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Can not connect to Client at %s error: %v \n", client_socket,err)
	}
	defer func() {
		connToClient.Close()
		fmt.Printf("Closed Connection To Client %s\n", client_socket)
	}()
	fmt.Printf("Connected To Client %s\n", client_socket)

	file_confirm_client:=upload.NewUploadServiceClient(connToClient)
	fmt.Print("Sending Notification To Client ....\n")
	
	// println("Sleeping")
	// time.Sleep(50 * time.Second)
	// // time.Sleep(10 * time.Second)
	// println("GoodMorning")

	res, err:=file_confirm_client.ConfirmUpload(context.Background(),&upload.ConfirmUploadRequest{
		FileName: fileName,
	})
	if err!=nil{
		fmt.Println("Failed to Send Notification to Client", err)
		return
	}

	response_status:=res.GetStatus()

	
	if response_status=="time_out"{
		fmt.Printf("File %s Confirmation is TimedOut So We will Drop it\n",fileName)
		//Remove File
		s.files_lookup_table.RemoveFile(fileName)

		//[TODO Extra Send Request to DataNode to Remove the FIle Uploaded]

		// Remove Client
		s.client_lookup_table.RemoveClient(fileName)

		return
	}
	if response_status=="wrong_file"{
		println("WRONG file Between expected by Node and ConfirmationSent [Syntax Error]")
		os.Exit(1)
		// [TODO] Wrong file Confirmation Handling

		return
	}
	
	//Update File as Confirmed
	s.files_lookup_table.ConfirmFile(fileName)
	// Remove Client
	s.client_lookup_table.RemoveClient(fileName)
	fmt.Print("Removed Client :D\n")
}

func main() {
	// Thread to listen to alive pings from data keepers
	fmt.Println("Hello From Master Node 😎")
	// Create Master Server
	master:=NewMasterServer()
	// (1) Register to the master node
	wg := sync.WaitGroup{}
	// add 2 goroutines to the wait group
	wg.Add(3)
	go func() {
		defer wg.Done()
		handleClient(&master)
	}()
	go func() {
		defer wg.Done()
		handleDataKeeper(&master)
	}()
	go func() {
		defer wg.Done()
		periodicCheckup(&master)
	}()

	// wait for all goroutines to finish
	wg.Wait()
}
