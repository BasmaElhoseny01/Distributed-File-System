package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"sync"

	// "time"

	"google.golang.org/grpc"

	tr "mp4-dfs/schema/file_transfer"
	// hb "mp4-dfs/schema/heart_beat"
	Reg "mp4-dfs/schema/register"
)

type nodeKeeperServer struct {
	tr.UnimplementedFileTransferServiceServer

	Ip string
	port string
}

func NewNodeKeeperServer(ip string, port string) *nodeKeeperServer {
	return &nodeKeeperServer{Ip: ip, port: port}
}

// FileTransfer Services rpc [client-streaming] RPC 
func (s *nodeKeeperServer) UploadFile(stream tr.FileTransferService_UploadFileServer) error {
	fmt.Println("Data: Received Uploading")

	// Receive Video Info
	req, err := stream.Recv()
	if err!=nil{
		fmt.Println("cannot receive file data",err)
		return err
	}

	fileName :=req.GetInfo().GetName()
	fmt.Println("received an upload-video request for file",fileName)

	// Receive Chunks
	video := bytes.Buffer{}
	videoSize:=0

	for {
		fmt.Println("Waiting to receive more data")

		req,err:=stream.Recv()
		if err == io.EOF{
			fmt.Println("No more Data")
			break
		}
		if err != nil {
			fmt.Println("cannot receive chunk data",err)
			return err
		}

		chunk:=req.GetChuckData()
		size := len(chunk)
		videoSize+=size

		fmt.Printf("received a chunk with size: %d\n", size)

		// Write the new chunk
		_,err=video.Write(chunk)
		if err!=nil{
			fmt.Println("cannot write chunk data",err)
			return err
		}
	}

	// Save To Disk
	err = writeVideoToDisk(fileName,video)
	if err !=nil{
		fmt.Println("ERRRRRRRRR")
	}

	res := &tr.UploadVideoResponse{
		Size: uint32(videoSize),
	}

	err = stream.SendAndClose(res)
	if err != nil {
		fmt.Printf("cannot send response: %v\n", err)
	}

	// Implement the RPC
	return nil
}

func writeVideoToDisk(fileName string,fileData bytes.Buffer) error{
	// 1. Create File
	filePath :=fileName
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("cannot create file at",filePath)
		return err
	}

	//2. Write to File
	_, err = fileData.WriteTo(file)
	if err != nil {
		fmt.Println("cannot write to file",err)
		return err
	}
	fmt.Printf("Saved %s at %s",fileName,filePath)
	return nil
}

// Ping Thread
func handlePing(connToMaster *grpc.ClientConn, id string) {
	// Register to HeartBeats Service
	// client := hb.NewHeartBeatServiceClient(connToMaster)

	// for {
	// 	_, err := client.AlivePing(context.Background(), &hb.AlivePingRequest{DataKeeperId: id})
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	// sleep for 1 seconds
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("Alive Ping Sent")
	// }
}

// Client Thread
func handleClient(ip string ,port string){
	fmt.Println("Handle Client")
	// listen to the port
	client_listener, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		fmt.Println(err)
	}
	defer client_listener.Close()

	// Define NodeKeeperServer
	data_keeper := NewNodeKeeperServer(ip, port)

	// define Data Keeper Server and register the service
	s := grpc.NewServer()

	// Register to FileTransfer Service [Server]
	tr.RegisterFileTransferServiceServer(s, data_keeper)

	if err := s.Serve(client_listener); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Handle Client Keeper finished")
}

var id string

func main() {
	fmt.Println("Hello From Data Node 📑")
	// TODO Fix That to be arguments
	ip:="localhost"
	port:="5003"
	
	// (1) Register to the master node
	masterAddress := "localhost:5002"

	connToMaster, err := grpc.Dial(masterAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer connToMaster.Close()

	// Register to New Node Registration Service
	client := Reg.NewDataKeeperRegisterServiceClient(connToMaster)

	// TODO (1) Register to the master node
	response, err := client.Register(context.Background(), &Reg.DataKeeperRegisterRequest{Ip: ip, Port: port})
	if err != nil {
		fmt.Println(err)
	}
	// print the response
	fmt.Println("Registered to the master node")
	id = response.GetDataKeeperId()
	fmt.Println("Data Keeper ID: ", id)

	wg := sync.WaitGroup{}
	// add 2 goroutines to the wait group
	wg.Add(2)
	go func() {
		defer wg.Done()
		handlePing(connToMaster,id)	
	}()
	go func() {
		defer wg.Done()
		handleClient(ip,port)
	}()

	wg.Wait()
}
