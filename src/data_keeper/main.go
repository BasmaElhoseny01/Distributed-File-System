package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"

	// 	"bytes"
	// 	"context"
	// 	"fmt"
	// 	"io"
	// 	"net"
	// 	"os"
	// 	"sync"
	// "time"
	// "context"

	"mp4-dfs/data_keeper/file_system_lookup"
	Reg "mp4-dfs/schema/register"
	utils "mp4-dfs/utils"

	download "mp4-dfs/schema/download"
	replicate "mp4-dfs/schema/replicate"
	upload "mp4-dfs/schema/upload"

	// "sync"
	// 	tr "mp4-dfs/schema/file_transfer"
	hb "mp4-dfs/schema/heart_beat"
	// 	fi "mp4-dfs/schema/finish_file_transfer"
)

type nodeKeeperServer struct {
	// tr.UnimplementedFileTransferServiceServer``
	upload.UnimplementedUploadServiceServer
	download.UnimplementedDownloadServiceServer
	replicate.UnimplementedReplicateServiceServer

	Id string
	Ip string
	file_port string
	replication_port string

	file_system_lookup_table file_system_lookup.FileSystemLookUpTable
}

func NewNodeKeeperServer(id string,ip string, file_port string ,replication_port string) *nodeKeeperServer {
	// Create or empty folder based on its existence
    if _, err := os.Stat(id); os.IsNotExist(err) {
        // Folder doesn't exist, create it
        if err := os.MkdirAll(id, os.ModePerm); err != nil {
            fmt.Printf("Error Creating Folder %v\n", err)
            os.Exit(1)
        }
        fmt.Printf("Folder %s created successfully\n", id)
    } else {
        // Folder exists, empty it
        if err := utils.EmptyFolder(id); err != nil {
            fmt.Printf("Error Emptying Folder %s: %v\n", id, err)
            os.Exit(1)
        }
        fmt.Printf("Folder %s emptied successfully\n", id)
    }

	return &nodeKeeperServer{
		Id:id, Ip: ip,file_port:file_port, replication_port:replication_port,
		file_system_lookup_table: file_system_lookup.NewFileSystemLookUpTable(),
	}
}

// ############################################################## RPCs #################################################################
// UploadFile rpc [client-streaming]
func (s *nodeKeeperServer) UploadFile(stream upload.UploadService_UploadFileServer) error{
	// Receive Video Info
	req, err := stream.Recv()
	if err!=nil{
		fmt.Println("Can not receive file data",err)
		return err
	}

	fileName:=req.GetFileInfo().GetFileName()
	fmt.Printf("Uploading %s .........\n",fileName)

	// Receive Chunks
	video := bytes.Buffer{}
	videoSize:=0

	for{
		req,err:=stream.Recv()
		if err == io.EOF{
			fmt.Println("Received EOF")
			break
		}
		if err != nil {
			fmt.Println("Can not receive chunk data",err)
			return err
		}

		chunk:=req.GetChuckData()
		size := len(chunk)
		videoSize+=size
		// fmt.Printf("Received a chunk with size: %d\n", size)
	
	
		// Write the new chunk
		_,err=video.Write(chunk)
		if err!=nil{
			fmt.Println("cannot write chunk data",err)
			return err
		}

	}
	// Save To Disk
	savePath:=s.Id+"/"+fileName
	err = writeVideoToDisk(savePath,video)
	if err !=nil{
		return err
	}

	// Add File to Files LookUpTable
	newFile:=file_system_lookup.NewFile(fileName,savePath)
	err=s.file_system_lookup_table.AddFile(&newFile)
	if err!=nil{
		fmt.Printf("Error When adding file %s to file system lookup Table at %s\n error %v\n",fileName,savePath,err)
	}
	fmt.Printf("New File added Successfully\n")
	fmt.Println(s.file_system_lookup_table.PrintFileInfo(fileName))

	
	//Send Final Response Close Connection With The Client
	err = stream.SendAndClose(&upload.UploadFileResponse{})
	if err != nil {
		fmt.Printf("Can not send Close response: %v\n", err)
		return err
	}
	fmt.Println("Connection With Client is Closed :D")

	//(2) Confirm To master the File Transfer
	// Establish Connection To Master
	masterAddress:=utils.GetMasterIP("node")
	connToMaster, err := grpc.Dial(masterAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Can not connect to Master at %s\n",masterAddress)
		return err
	}
	fmt.Println("Connected To Master at", masterAddress)

	// Register to Upload File Service with master as server
	upload_file_client :=upload.NewUploadServiceClient(connToMaster)
	fmt.Printf("Sending File %s Upload Confirm To Master .....\n", fileName)

	_,err=upload_file_client.NotifyMaster(context.Background(),&upload.NotifyMasterRequest{
		NodeId: s.Id,
		FileName: fileName,
	})
	if err!=nil{
		fmt.Printf("Failed to Notify Master with uploading file %s %v\n",fileName,err)
		return err
	}
	fmt.Printf("Sent File %s Upload Confirm To Master\n",fileName)
	connToMaster.Close()
	fmt.Printf("Closed Connection to Master at %s\n",masterAddress)

	return nil
}

// DownloadFile rpc [server-streaming]
func (s *nodeKeeperServer) Download(req *download.DownloadRequest, stream download.DownloadService_DownloadServer) error {
	fileName:=req.GetFileName()
	offset := req.GetOffset()
	skip := req.GetSkip()
	fmt.Printf("Downloading %s .........\n",fileName)

	// Get Path of this file in the node system
	filePath:=s.file_system_lookup_table.GetFilePath(fileName)

	// Open File
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Cannot open Video File at [%s] got error: %v\b", filePath,err)
		return err
	}
	defer file.Close()

	// Read File
	chunk := make([]byte, 1024*1024) // 1MB
	// Skip Offset where offset is the id and bytes to skip is offset*1024*1024
	_, err = file.Seek(int64(offset*1024*1024), 0)
	if err != nil {
		fmt.Println("Cannot seek file",err)
		return err
	}
	id := offset
	for {
		n, err := file.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Cannot read file",err)
			return err
		}

		// Send Chunk
		err = stream.Send(&download.DownloadResponse{Data: &download.DownloadResponse_Chunk{
			// Send the chunk of data and chunk id
			Chunk: &download.ChunkData{
				Data:    chunk[:n],
				ChunkId: int64(id),
			},
		}})
		if err != nil {
			fmt.Println("Cannot send chunk",err)
			return err
		}
		// Skip the bytes
		if skip > 0 {
			_, err = file.Seek(int64(skip*1024*1024), 1)
			if err != nil {
				fmt.Println("Cannot seek file",
				err)
				return err
			}
		}
		id += skip+1
	}

	fmt.Printf("Done Downloading %s \n",fileName)

	return nil
}

// Replication RPCs
// Master Notify To Data Node to Copy
func (s *nodeKeeperServer) NotifyToCopy (ctx context.Context, in *replicate.NotifyToCopyRequest) (*replicate.NotifyToCopyResponse,error){
	file_name:=in.GetFileName()
	srcAddress:=in.GetSrcAddress()
	fmt.Printf("I get Notified to Copy File %s from Node %s\n",file_name,srcAddress)

	go func() {
		s.handleCopy(file_name,srcAddress)
	}()


	// Response on Master
	return &replicate.NotifyToCopyResponse{Status: "ok"},nil
}

// Coping rpc [server-streaming]
func (s *nodeKeeperServer) Copying(req *replicate.CopyingRequest, stream replicate.ReplicateService_CopyingServer) error {
	fileName:=req.GetFileName()

	// Get Path of this file in the node system
	filePath:=s.file_system_lookup_table.GetFilePath(fileName)

	// Open File
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Cannot open Video File at [%s] got error: %v\b", filePath,err)
		return err
	}
	defer file.Close()

	//3. Send MetaData For Video File
	err = stream.Send(&replicate.CopyingResponse{
		Data: &replicate.CopyingResponse_FileInfo{FileInfo: &replicate.FileInfo{FileName: fileName}},
	})
	if err != nil {
		fmt.Println("cannot send file info to server: ", err, stream.RecvMsg(nil))
		return err
	}

	// Read File
	chunk := make([]byte, 1024*1024) // 1MB
	for {
		n, err := file.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Cannot read file",err)
			return err
		}
		// Send Chunk
		err = stream.Send(&replicate.CopyingResponse{Data: &replicate.CopyingResponse_ChuckData{ChuckData: chunk[:n]}})
		if err != nil {
			fmt.Println("Cannot send chunk",err)
			return err
		}
	}
	fmt.Println("Sent File Successfully")
	return nil
}


// #################################################################### Utils ########################################################
func listenOnPort(server *grpc.Server, ip string ,port string) {
	socket:=ip+":"+port
    // Listen for incoming connections on the specified port
	client_listener, err := net.Listen("tcp",socket )
	fmt.Printf("Listening to Socket %s\n",socket)

	if err != nil {
		fmt.Println(err)
	}
	defer client_listener.Close()

	if err := server.Serve(client_listener); err != nil {
		fmt.Println(err)
	}
}

func writeVideoToDisk(filePath string,fileData bytes.Buffer) error{
	// 1. Create File
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Can not create file at",filePath)
		return err
	}

	//2. Write to File
	_, err = fileData.WriteTo(file)
	if err != nil {
		fmt.Println("Can not write to file",err)
		return err
	}
	fmt.Printf("Saved File at %s\n",filePath)

	return nil
}

// #################################################################### GO Routines ##################################################
// Ping Thread
func handlePing() {
	fmt.Println("Handle Ping")

	// 1. Connect To Master [Ping Port]
	masterPingAddress:=utils.GetMasterIP("ping")	
	// Dial
	connToPingMaster, err := grpc.Dial(masterPingAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected To Master at", masterPingAddress,"[Ping]")

	// 2. Register to HeartBeats Service
	client := hb.NewHeartBeatServiceClient(connToPingMaster)

	for {
		_, err := client.AlivePing(context.Background(), &hb.AlivePingRequest{DataKeeperId: id})
		if err != nil {
			fmt.Println(err)
		}
		// sleep for 1 seconds
		time.Sleep(1 * time.Second)
		// fmt.Println("Alive Ping Sent")
	}
}

// Client Thread
func handleClient(data_keeper *nodeKeeperServer,ip string ,port string){
	fmt.Println("Handle Client")

	// define Data Keeper Server and register the service
	s := grpc.NewServer()

	// Register to Upload File Service [Server]
	upload.RegisterUploadServiceServer(s, data_keeper)
	
	// Register to Download File Service [Server]
	download.RegisterDownloadServiceServer(s, data_keeper)

	// Listen to incoming requests on that port :D
	listenOnPort(s,ip,port)

	// Keep the main goroutine running
	select {}
}

// Replicate Thread
func handleReplicate(data_keeper *nodeKeeperServer,ip string ,port string){
	fmt.Println("Handle Replicate")

	// define Data Keeper Server and register the service
	s := grpc.NewServer()

	// Register to Replication Service [Server]
	replicate.RegisterReplicateServiceServer(s, data_keeper)

	// Listen to incoming requests on that port :D
	listenOnPort(s,ip,port)

	// Keep the main goroutine running
	select {}
}

// Copying Thread
func (s *nodeKeeperServer) handleCopy(file_name string, src_address string){
	
	//1. Establish connection with srcAddress
	connToSrc,err:= grpc.Dial(src_address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Cannot connect to Data Node at", src_address)
		return 
	}
	fmt.Printf("Connected To Src Data Node %s\n", src_address)

	// 2. Register as Client
	replicateClient:=replicate.NewReplicateServiceClient(connToSrc)
	stream,err :=replicateClient.Copying(context.Background(),&replicate.CopyingRequest{
		FileName: file_name,
	})
	if err != nil {
		fmt.Println("Cannot Copying file", err)
		return
	}

	if break_node{
		fmt.Println("Shutting Down ...")
		os.Exit(0)
	}

	// create file
	// Save To Disk
	savePath:=id+"/"+file_name
	file, err := os.Create(savePath)
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Cannot receive chunk", err)
			return
		}
		_, err = file.Write(chunk.GetChuckData())
		if err != nil {
			fmt.Println("Cannot write chunk to file", err)
			return
		}
	}
	file.Close()
	fmt.Println("File Copying successfully")
	connToSrc.Close()

	// Add File to Files LookUpTable
	newFile:=file_system_lookup.NewFile(file_name,savePath)
	err=s.file_system_lookup_table.AddFile(&newFile)
	if err!=nil{
		fmt.Printf("Error When adding file %s to file system lookup Table at %s\n error %v\n",file_name,savePath,err)
	}
	fmt.Printf("New File added Successfully\n")
	fmt.Println(s.file_system_lookup_table.PrintFileInfo(file_name))	

	// 4. Send Ack to Master
	masterIP:=utils.GetMasterIP("node")
	connToMaster, err := grpc.Dial(masterIP, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Can not connect to Master at %s, Error %s\n", masterIP,err)
	}

	fmt.Printf("Connected To Src %s\n", masterIP)
	replicateClient=replicate.NewReplicateServiceClient(connToMaster)

	fmt.Printf("Sending Ack Master\n")
	_ ,err=replicateClient.ConfirmCopy(context.Background(),&replicate.ConfirmCopyRequest{
		FileName: file_name,
		NodeId: s.Id,
	})
	if err!=nil{
		fmt.Println("Failed to sending ACK To Master", err)
	}
	fmt.Printf("Replication Ack Sent to Master\n")

	connToMaster.Close()
}

var id string

func GetNodeSockets() (node_ip string, file_service_port_no string ,replication_service_port_no string,breakEnabled bool) {
	// Define flags
	var breakFlag bool
	flag.BoolVar(&breakFlag, "break", false, "Enable break")

	// Parse flags
	flag.Parse()

	if breakFlag{
		// Remove the element at index 2
		os.Args = append(os.Args[:1], os.Args[2:]...)
	}

	// Take The port Nos from Command Line
	if(len(os.Args)<=1){
		// go run ./data_keeper/main.go --> MyIP + Empty Port [Done]

		// Get IP address using GetMyIP function if not provided in arguments
		ip := utils.GetMyIp()
		if ip == nil {
			fmt.Println("Failed to retrieve IP address.")
			os.Exit(1)
		}

		//No ports are passed then get empty one
		file_service_port,err:=utils.GetEmptyPort(ip.String(),"-1")
		if err!=nil{
			fmt.Printf("Error When Getting Empty Port for the file_service_port Error:%v",err)
			os.Exit(1)
		}

		//No ports are passed then get empty one
		replication_service_port,err:=utils.GetEmptyPort(ip.String(),file_service_port)
		if err!=nil{
			fmt.Printf("Error When Getting Empty Port for the replication_service_port Error:%v",err)
			os.Exit(1)
		}

		return ip.String(),file_service_port,replication_service_port,breakFlag
	}


	// Check if enough arguments are provided
	  if len(os.Args) > 5 {
        fmt.Println("Usage: data_node [-break] [<your_ip>] [<file_service_port>] [<replication_service_port>]")
        os.Exit(1)
    }

	// If the first argument is an IP address, parse it
    ip := net.ParseIP(os.Args[1])
	if ip == nil{
		ipAddr, err := net.ResolveIPAddr("ip", os.Args[1])
		if err==nil{
			ip=ipAddr.IP
		}
	}

	// [IP]
    if ip != nil {
        // If the first argument is an IP address, shift arguments to the right
        os.Args = append(os.Args[:1], os.Args[2:]...)
    } else {
        // Get IP address using GetMyIP function if not provided in arguments
        ip = utils.GetMyIp()
        if ip == nil {
            fmt.Println("Failed to retrieve IP address.")
			os.Exit(1)
        }
	}

	// [PORT]
	if len(os.Args)>2{
		// Then a ports are passed
		// Parse client port
		file_service_port, err := strconv.Atoi(os.Args[len(os.Args)-2])
		if err != nil || file_service_port <= 0 || file_service_port > 65535 {
			fmt.Println("Invalid file service port:", os.Args[len(os.Args)-2])
			os.Exit(1)
		}

		// Check if client port is reachable
		if !utils.IsPortOpen(ip.String(), file_service_port) {
			fmt.Printf("File service port %d is not reachable\n", file_service_port)
			os.Exit(1)
		}
		file_service_port_no=strconv.Itoa(file_service_port)


		// Parse node port
		replication_service_port, err := strconv.Atoi(os.Args[len(os.Args)-1])
		if err != nil || replication_service_port <= 0 || replication_service_port > 65535 {
			fmt.Println("Invalid replication service port:", os.Args[len(os.Args)-1])
			os.Exit(1)
		}

		// Check if node port is reachable
		if !utils.IsPortOpen(ip.String(), replication_service_port) {
		fmt.Printf("Replication service port %d is not reachable\n", replication_service_port)
		os.Exit(1)
		}
		replication_service_port_no=strconv.Itoa(replication_service_port)
	}else{
		//No ports are passed then get empty one
		file_service_port,err:=utils.GetEmptyPort(ip.String(),"-1")
		if err!=nil{
			fmt.Printf("Error When Getting Empty Port for the file_service_port Error:%v",err)
			os.Exit(1)
		}
		file_service_port_no=file_service_port

		//No ports are passed then get empty one
		replication_service_port,err:=utils.GetEmptyPort(ip.String(),file_service_port)
		if err!=nil{
			fmt.Printf("Error When Getting Empty Port for the replication_service_port Error:%v",err)
			os.Exit(1)
		}
		replication_service_port_no=replication_service_port

	}

	return ip.String(),file_service_port_no,replication_service_port_no,breakFlag
}

var break_node bool=false

func main() {
	fmt.Println("Hello From Data Node 📑")

	// 1. Get Ip & Ports
	ip,file_service_port,replication_service_port,breakEnabled:=GetNodeSockets()
	// ip:="127.0.0.1"
	// file_service_port="8080"
	// replication_service_port="8085"
	//breakEnabled=false

	// Now you can use ip and ports in your program
	fmt.Println("IP address:", ip)
	fmt.Println("File Service port:", file_service_port)
	fmt.Println("Replication Service port:", replication_service_port)
	if breakEnabled{
		break_node=true
		fmt.Println("You asked me to Break",breakEnabled)
	}

	// 2. Connect To Master
	masterAddress:=utils.GetMasterIP("node")
	// Dial
	connToMaster, err := grpc.Dial(masterAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected To Master at", masterAddress,"[Data]")

	// 3. Register to the master node
	// Register to New Node Registration Service
	client := Reg.NewDataKeeperRegisterServiceClient(connToMaster)

	response, err := client.Register(context.Background(), &Reg.DataKeeperRegisterRequest{Ip: ip, FilePort: file_service_port, ReplicationPort: replication_service_port})
	if err != nil {
		fmt.Println("Can't Register The Data Node to Master",err)
		os.Exit(0)
	}
	id = response.GetDataKeeperId()
	fmt.Printf("Registered To Master With ID %s\n", id)
	
	// Close Connection with the master for port of data [Not Ping]
	connToMaster.Close()

	// Create DataNode Server
	data_keeper_server := NewNodeKeeperServer(id,ip,file_service_port,replication_service_port)

	wg := sync.WaitGroup{}
	// add 2 goroutines to the wait group
	wg.Add(2)
	go func() {
		defer wg.Done()
		handlePing()	
	}()
	go func() {
		defer wg.Done()
		handleClient(data_keeper_server,ip,file_service_port)
	}()
	go func() {
		defer wg.Done()
		handleReplicate(data_keeper_server,ip,replication_service_port)
	}()
	wg.Wait()
}


// go run .\data_keeper\main.go --> MyIp + Empty Port + Empty Port 
// go run .\data_keeper\main.go 127.0.0.1 8080 8085 
// go run .\data_keeper\main.go 8080 8085 
// go run .\data_keeper\main.go 127.0.0.1 
// go run .\data_keeper\main.go  127.0.0.1 9040 9806
// go run .\data_keeper\main.go  127.0.0.1 9041 9807


// go run .\data_keeper\main.go -break --> MyIp + Empty Port + Empty Port
// go run .\data_keeper\main.go -break 127.0.0.1 9040 9806
// go run .\data_keeper\main.go -break 8080 8085
// go run .\data_keeper\main.go -break 127.0.0.1