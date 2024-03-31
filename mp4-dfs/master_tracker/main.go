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

	cf "mp4-dfs/schema/confirm_file_transfer"
	download "mp4-dfs/schema/download"
	hb "mp4-dfs/schema/heart_beat"
	reg "mp4-dfs/schema/register"
	replicate "mp4-dfs/schema/replicate"
	upload "mp4-dfs/schema/upload"
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


// ################################################################ RPCs ################################################################
// DataKeepersNodes Registration Services rpc
func (s *masterServer) Register(ctx context.Context, in *reg.DataKeeperRegisterRequest) (*reg.DataKeeperRegisterResponse, error) {
	Ip:=in.GetIp()
	file_port:=in.GetFilePort()
	replication_port:=in.GetReplicationPort()

	// Add the data node to the lookup table
	new_data_node:=data_lookup.NewDataNode(Ip,file_port,replication_port)
	node_id, err := s.data_node_lookup_table.AddDataNode(&new_data_node)
	if err!=nil{
		fmt.Printf("Error When adding new Data Node with ID: %s",node_id)
		return  &reg.DataKeeperRegisterResponse{},err
	}
	fmt.Printf("New Data Node added Successfully\n")
	fmt.Println(s.data_node_lookup_table.PrintDataNodeInfo(node_id))

	return &reg.DataKeeperRegisterResponse{DataKeeperId: node_id}, nil
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
	node_id,err:=s.data_node_lookup_table.GetLeastLoadedNode()
	if err != nil {
		fmt.Printf("Can not Get Least Loaded Node DataNode  %v\n",err)
		return  &upload.RequestUploadResponse{},err
	}

	// Get File Service Socket for Node
	ip,port,err:=s.data_node_lookup_table.GetNodeFileServiceAddress(node_id)
	if err != nil {
		fmt.Printf("Can not Get DataNode [%s] Port %v\n",node_id,err)
		return  &upload.RequestUploadResponse{},err
	}
	node_socket:=ip+":"+port

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


// DownloadService RpcS
// Get Server RPC
// BASMA
func (s *masterServer) GetServer(ctx context.Context, in *download.DownloadRequest) (*download.DownloadServerResponse, error) {
	file_name:=in.GetFileName()
	fmt.Println("Received Download Request",file_name)
	// check if file already exist
	exists,node_1,node_2,node_3 :=s.files_lookup_table.GetFile(file_name)
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
		Ip1,Port1,_ =s.data_node_lookup_table.GetNodeFileServiceAddress(node_1)
	}
	if node_2!= "-1"{
		Ip2,Port2,_=s.data_node_lookup_table.GetNodeFileServiceAddress(node_2)
	}
	if node_3!= "-1"{
		Ip3,Port3,_ =s.data_node_lookup_table.GetNodeFileServiceAddress(node_3)
	}
	// create list of servers which contains ip and port
	servers := []*download.Server{
	}
	if Ip1!=""{
		servers = append(servers, &download.Server{
			Ip:   Ip1,
			Port: Port1,
		})
	}
	if Ip2!=""{
		servers = append(servers, &download.Server{
			Ip:   Ip2,
			Port: Port2,
		})
	}
	if Ip3!=""{
		servers = append(servers, &download.Server{
			Ip:   Ip3,
			Port: Port3,
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

// Heart Beat Services PRCs
// HeartBeat Registration Services rpc
func (s *masterServer) AlivePing(ctx context.Context, in *hb.AlivePingRequest) (*hb.AlivePingResponse, error) {
	node_id:=in.GetDataKeeperId()
	stamp,err:=s.data_node_lookup_table.UpdateNodeTimeStamp(node_id)

	if err == nil {
		fmt.Printf("Data Node '%s' Ping Time stamp Updated with %s \n", node_id,stamp.Format("2006-01-02 15:04:05"))
	}
	return &hb.AlivePingResponse{},nil
}

//################################################################ Go Routines ################################################################
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
	fmt.Printf("Listening to Data Keeper at Socket: %s [Data]\n",masterAddress)

	// define our master server and register the service
	s := grpc.NewServer()

	// Register in New Node Registration Service
	reg.RegisterDataKeeperRegisterServiceServer(s,master)

	// Register to UploadService
	upload.RegisterUploadServiceServer(s,master)

	// Register to DownloadService
	download.RegisterDownloadServiceServer(s,master)
	
	if err := s.Serve(dataKeeper_listener); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Handle Data Keeper finished")
}


func handleDataKeeperPing(master *masterServer) {
	// listen to the port
	masterAddress:=utils.GetMasterIP("ping")
	dataKeeper_listener, err := net.Listen("tcp", masterAddress)
	if err != nil {
		fmt.Println(err)
	}
	defer dataKeeper_listener.Close()
	fmt.Printf("Listening to Data Keeper at Socket: %s [Ping]\n",masterAddress)

	// define our master server and register the service
	s := grpc.NewServer()

	// Register in HeartBeat Service
	hb.RegisterHeartBeatServiceServer(s,master)
	
	if err := s.Serve(dataKeeper_listener); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Handle Data Keeper Ping finished")
}

func checkIdleNodes(master *masterServer){
	for{
		//1. Check Ideal 
		print("Check Ideal Nodes....\n")
		master.data_node_lookup_table.CheckPingStatus()

		// Sleep for 5 seconds before the next check
		time.Sleep(5 * time.Second)
	}

}

func checkUnConfirmedFiles(master *masterServer){
	for{
		// 2.Sent Notifications to Clients
		println("Checking UnConfirmed Files...\n")
		unconfirmedFiles:=master.files_lookup_table.CheckUnConfirmedFiles()
		println(unconfirmedFiles)
		for _, file := range unconfirmedFiles{
			master.sendClientConfirm(file)

		}

		// Sleep for 5 seconds before the next check
		time.Sleep(5 * time.Second)
	}

}

func checkReplication(master *masterServer) {   
	for {
		println("Checking UnReplicated Files...\n")

		// Iterate through distinct file instances
		// Getting non replicated files
		files := master.files_lookup_table.CheckUnReplicatedFiles()
		for _, file := range files {
			sourceMachines := master.files_lookup_table.GetFileSourceMachines(file)
	
			// Get first Source
			srcId:=sourceMachines[0]
			srcIP,srcPort,_:=master.data_node_lookup_table.GetNodeReplicationServiceAddress(srcId)
			srcAddress:=srcIP+":"+srcPort
			// replica_count:=len(sourceMachines)
			// for{
			// 	if replica_count==3{
			// 		break
			// 	}
				dstId,_ := master.data_node_lookup_table.GetCopyDestination(sourceMachines)
				if dstId !=""{
					// 1.Send Dst IP to Src
					// Append
					dstIP,dstPort,_:=master.data_node_lookup_table.GetNodeReplicationServiceAddress(dstId)
					dstAddress:=dstIP+":"+dstPort
					connToDst, err := grpc.Dial(dstAddress, grpc.WithInsecure())
					if err != nil {
						fmt.Println(err)
						fmt.Printf("Can not connect to Dst at %s, Error %s\n", dstAddress,err)
					}
	
					fmt.Printf("Connected To Dst %s\n", dstAddress)
	
					//2. Register as Client to Service replicate File offered by the datanode
					replicateClient := replicate.NewReplicateServiceClient(connToDst)
	
					//3- sending request
					fmt.Printf("Sending copy request to data node\n")
					_ ,err=replicateClient.NotifyToCopy(context.Background(),&replicate.NotifyToCopyRequest{
						FileName: file,
						SrcAddress: srcAddress,
					})
					if err!=nil{
						fmt.Println("Failed to Notify Dst Node for Replication", err)
					}
					connToDst.Close()
		
				}else{
					println("No available Data Nodes to replicate the file.\n")
				}
	
						
				// 3. Check For Replicas
		}
		
		// [FIX] Sleep for 10 seconds before the next check
		time.Sleep(2 * time.Second)
    }
}

// ###################################################### Utils ##########################################################
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
		// Checked this Mechanism and found best way to handle timed_out is to keep the file.
		fmt.Printf("File %s Confirmation is TimedOut So We Only Remove Client From Table But The File is Kept\n",fileName)
		//Remove File
		// s.files_lookup_table.RemoveFile(fileName)

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
	// [FIX] Serve MultiCalls
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
		// Thread for handling DataNodes for Registration and Requests of Files
		defer wg.Done()
		handleDataKeeper(&master)
	}()
	go func() {
		// Thread for handling Pings from Data Node
		defer wg.Done()
		handleDataKeeperPing(&master)
	}()
	go func() {
		defer wg.Done()
		checkIdleNodes(&master)
	}()
	go func() {
		defer wg.Done()
		checkUnConfirmedFiles(&master)
	}()
	go func() {
		defer wg.Done()
		checkReplication(&master)
	}()

	// wait for all goroutines to finish
	wg.Wait()
}

// go run .\master_tracker\main.go