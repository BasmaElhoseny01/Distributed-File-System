package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"

	// 	"bytes"
	// 	"context"
	// 	"fmt"
	// 	"io"
	// 	"net"
	// 	"os"
	// 	"sync"
	// 	// "time"
	// "context"

	Reg "mp4-dfs/schema/register"
	utils "mp4-dfs/utils"
	// "sync"
	// 	tr "mp4-dfs/schema/file_transfer"
	// 	// hb "mp4-dfs/schema/heart_beat"
	// 	fi "mp4-dfs/schema/finish_file_transfer"
	// 	// download "mp4-dfs/schema/download"
)

// type nodeKeeperServer struct {
// 	tr.UnimplementedFileTransferServiceServer

// 	Id string
// 	Ip string
// 	port string
// }

// func NewNodeKeeperServer(id string,ip string, port string) *nodeKeeperServer {
// 	return &nodeKeeperServer{Id:id, Ip: ip, port: port}
// }

// // FileTransfer Services rpc [client-streaming] RPC
// func (s *nodeKeeperServer) UploadFile(stream tr.FileTransferService_UploadFileServer) error {
// 	fmt.Println("Data: Received Uploading")

// 	// Receive Video Info
// 	req, err := stream.Recv()
// 	if err!=nil{
// 		fmt.Println("cannot receive file data",err)
// 		return err
// 	}

// 	fileName :=req.GetInfo().GetName()
// 	fmt.Println("received an upload-video request for file",fileName)

// 	// Receive Chunks
// 	video := bytes.Buffer{}
// 	videoSize:=0

// 	for {
// 		fmt.Println("Waiting to receive more data")

// 		req,err:=stream.Recv()
// 		if err == io.EOF{
// 			fmt.Println("No more Data")
// 			break
// 		}
// 		if err != nil {
// 			fmt.Println("cannot receive chunk data",err)
// 			return err
// 		}

// 		chunk:=req.GetChuckData()
// 		size := len(chunk)
// 		videoSize+=size

// 		fmt.Printf("received a chunk with size: %d\n", size)

// 		// Write the new chunk
// 		_,err=video.Write(chunk)
// 		if err!=nil{
// 			fmt.Println("cannot write chunk data",err)
// 			return err
// 		}
// 	}

// 	// Save To Disk
// 	savePath:=s.Id+"/"+fileName
// 	err = writeVideoToDisk(savePath,video)
// 	if err !=nil{
// 		return err
// 	}

// 	res := &tr.UploadVideoResponse{
// 		Size: uint32(videoSize),
// 	}

// 	err = stream.SendAndClose(res)
// 	if err != nil {
// 		fmt.Printf("cannot send response: %v\n", err)
// 	}

// 	// (2) Confirm To master the File Transfer
// 	handleConfirmToMaster(s.Id,fileName,savePath)

// 	return nil
// }
// func handleConfirmToMaster(data_node_id string,file_name string, file_path string){
// 	// [TODO] Fix This Replication to connection
// 	//Establish Connection to Master Node
// 	masterAddress := "localhost:5002"

// 	connToMaster, err := grpc.Dial(masterAddress, grpc.WithInsecure())
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer connToMaster.Close()

// 	// Register To Confirm File Transfer Service
// 	finish_file_transfer_client :=fi.NewFinishFileTransferServiceClient(connToMaster)

// 	_,err = finish_file_transfer_client.FinishFileUpload(context.Background(),&fi.FinishFileUploadRequest{
// 		DataNodeId: data_node_id,
// 		FileName: file_name,
// 		FilePath: file_path,
// 	})
// 	if err!=nil{
// 		fmt.Print("Failed to Send Finish Upload File Request To Master:",err)
// 	}
// }

// func writeVideoToDisk(filePath string,fileData bytes.Buffer) error{

// 	//Create Folder
// 	_, err := os.Stat(id)
// 	if os.IsNotExist(err) {
//         // Folder doesn't exist, create it
//         err := os.MkdirAll(id, os.ModePerm)
//         if err != nil {
//             return err
//         }
//         fmt.Printf("Folder %s created successfully\n",id)
// 	}
// 	// 1. Create File
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		fmt.Println("cannot create file at",filePath)
// 		return err
// 	}

// 	//2. Write to File
// 	_, err = fileData.WriteTo(file)
// 	if err != nil {
// 		fmt.Println("cannot write to file",err)
// 		return err
// 	}
// 	fmt.Printf("Saved at %s",filePath)
// 	return nil
// }

// // Ping Thread
// func handlePing(connToMaster *grpc.ClientConn) {
// 	// Register to HeartBeats Service
// 	// client := hb.NewHeartBeatServiceClient(connToMaster)

// 	// for {
// 	// 	_, err := client.AlivePing(context.Background(), &hb.AlivePingRequest{DataKeeperId: id})
// 	// 	if err != nil {
// 	// 		fmt.Println(err)
// 	// 	}
// 	// 	// sleep for 1 seconds
// 	// 	time.Sleep(1 * time.Second)
// 	// 	fmt.Println("Alive Ping Sent")
// 	// }
// }

// // Client Thread
// func handleClient(ip string ,port string){
// 	fmt.Println("Handle Client")
// 	// listen to the port
// 	client_listener, err := net.Listen("tcp", ip+":"+port)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer client_listener.Close()

// 	// Define NodeKeeperServer
// 	data_keeper := NewNodeKeeperServer(id,ip, port)

// 	// define Data Keeper Server and register the service
// 	s := grpc.NewServer()

// 	// Register to FileTransfer Service [Server]
// 	tr.RegisterFileTransferServiceServer(s, data_keeper)

// 	if err := s.Serve(client_listener); err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println("Handle Client Keeper finished")
// }

// var id string

func GetNodeSockets() (node_ip string, node_ports []string) {
	// Take The port Nos from Command Line
	if len(os.Args) < 2 {
        fmt.Println("Usage: data_node [<your_ip>] <port1> [<port2> ...]")
        return
    }

	// If the first argument is an IP address, parse it
    ip := net.ParseIP(os.Args[1])
	if ip == nil{
		ipAddr, err := net.ResolveIPAddr("ip", os.Args[1])
		if err==nil{
			ip=ipAddr.IP
		}
	}
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

	var ports []int

    // Parse port numbers from the command line
    for i := 1; i < len(os.Args); i++ {
        port, err := strconv.Atoi(os.Args[i])
        if err != nil {
            fmt.Println("Invalid port number:", os.Args[i])
        }else {
			ports = append(ports, port)
		}
    }
	var reachable_ports []string
	// Check if the IP address and ports are reachable
    for _, port := range ports {
        if !utils.IsPortOpen(ip.String(), port) {
            fmt.Printf("Port %d is not reachable\n", port)
        }else{
			reachable_ports=append(reachable_ports, strconv.Itoa(port))
		}
	}

	if len(ports) == 0 {
		fmt.Println("At least one port is required.")
		os.Exit(1)
    }
	return ip.String(),reachable_ports
}

func main() {
	fmt.Println("Hello From Data Node 📑")

	// 1. Get Ip & Ports
	// ip,ports:=GetNodeSockets()
	ip:="127.0.0.1"
	ports:=[]string{"8080", "8081"}

	// Now you can use ip and ports in your program
	fmt.Println("IP address:", ip)
	fmt.Println("Ports:", ports)

	// 2. Connect To Master
	masterAddress:=utils.GetMasterIP("node")
	fmt.Println("Master Node Socket:", masterAddress)

	connToMaster, err := grpc.Dial(masterAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer connToMaster.Close()

	// 3. Register to the master node
	// Register to New Node Registration Service
	client := Reg.NewDataKeeperRegisterServiceClient(connToMaster)

	response, err := client.Register(context.Background(), &Reg.DataKeeperRegisterRequest{Ip: ip,Ports: ports})
	if err != nil {
		fmt.Println("Can't Register The Data Node to Master",err)
		os.Exit(0)
	}
	// print the response
	fmt.Println("Registered to the master node")
	id := response.GetDataKeeperId()
	fmt.Println("Data Keeper ID: ", id)

	// wg := sync.WaitGroup{}
	// // add 2 goroutines to the wait group
	// wg.Add(2)
	// go func() {
	// 	defer wg.Done()
	// 	handlePing(connToMaster)	
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	handleClient(ip,port)
	// }()

	// wg.Wait()
}


// go run .\data_keeper\main.go 127.0.0.1 8080 52001 -8