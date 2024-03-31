package main

import (
	"bytes"
	"context"
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

	Reg "mp4-dfs/schema/register"
	utils "mp4-dfs/utils"

	download "mp4-dfs/schema/download"
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
	Id string
	Ip string
	ports []string
}

func NewNodeKeeperServer(id string,ip string, ports[]string) *nodeKeeperServer {
	return &nodeKeeperServer{Id:id, Ip: ip,ports:ports }
}

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
	
	
	//Send Final Response Close Connection With The Client
	err = stream.SendAndClose(&upload.UploadFileResponse{})
	if err != nil {
		fmt.Printf("Can not send Close response: %v\n", err)
		return err
	}
	fmt.Println("Connection With Client is Closed :D")

	//(2) [TODO]Confirm To master the File Transfer
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
		FilePath: savePath,
	})
	if err!=nil{
		fmt.Printf("Failed to Notify Master with uploading file %s %v\n",fileName,err)
		// [TODO] Delete the file because saving is useless
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
	// Open File
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Cannot open Video File at [%s] got error: %v\b", fileName,err)
		return err
	}
	defer file.Close()

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
		err = stream.Send(&download.DownloadResponse{Data: &download.DownloadResponse_ChuckData{ChuckData: chunk[:n]}})
		if err != nil {
			fmt.Println("Cannot send chunk",err)
			return err
		}
	}
	return nil

}
func writeVideoToDisk(filePath string,fileData bytes.Buffer) error{

	//Create Folder
	_, err := os.Stat(id)
	if os.IsNotExist(err) {
        // Folder doesn't exist, create it
        err := os.MkdirAll(id, os.ModePerm)
        if err != nil {
            return err
        }
        fmt.Printf("Folder %s created successfully\n",id)
	}
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

// Client Thread
func handleClient(ip string ,ports []string){
	fmt.Println("Handle Client")
	
	// Define NodeKeeperServer
	data_keeper := NewNodeKeeperServer(id,ip,ports)

	// define Data Keeper Server and register the service
	s := grpc.NewServer()

	// Register to Upload File Service [Server]
	upload.RegisterUploadServiceServer(s, data_keeper)
	
	// Register to Download File Service [Server]
	download.RegisterDownloadServiceServer(s, data_keeper)
	// Loop through each port and start a listener
	for _, port := range ports {
        go listenOnPort(s,ip,port)
    }

	// Keep the main goroutine running
	select {}
}

var id string

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
	ports:=[]string{"8080"}

	// Now you can use ip and ports in your program
	fmt.Println("IP address:", ip)
	fmt.Println("Ports:", ports)

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

	response, err := client.Register(context.Background(), &Reg.DataKeeperRegisterRequest{Ip: ip,Ports: ports})
	if err != nil {
		fmt.Println("Can't Register The Data Node to Master",err)
		os.Exit(0)
	}
	id = response.GetDataKeeperId()
	fmt.Printf("Registered To Master With ID %s\n", id)

	// Close Connection with the master for port of data [Not Ping]
	connToMaster.Close()

	wg := sync.WaitGroup{}
	// add 2 goroutines to the wait group
	wg.Add(2)
	go func() {
		defer wg.Done()
		handlePing()	
	}()
	go func() {
		defer wg.Done()
		handleClient(ip,ports)
	}()

	wg.Wait()
}


// go run .\data_keeper\main.go 127.0.0.1 8080 -8