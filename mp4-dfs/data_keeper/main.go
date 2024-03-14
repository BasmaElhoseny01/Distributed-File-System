package main

import (
	"context"
	"fmt"
	"net"
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

	for {
		// req,err = stream.Recv()
		// fmt.r
		// if err ==io.EOF{
		// 	//End Of Stream :D
		// }
	}

	// Implement thr RPC
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
