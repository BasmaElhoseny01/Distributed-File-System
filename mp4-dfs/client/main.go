package main

import (
	"bufio"
	"context"
	"fmt"
	"io"

	// "net"
	"os"
	"path/filepath"

	"google.golang.org/grpc"

	upload "mp4-dfs/schema/upload"
	utils "mp4-dfs/utils"
	// download "mp4-dfs/schema/download"
)

type clientNode struct{
	upload.UnimplementedUploadServiceServer
	socket string
}

func newClientNode(socket_no string) clientNode{
	return clientNode{
		socket:socket_no,
	}
}

// ConfirmUpload rpc
func (c *clientNode) ConfirmUpload(ctx context.Context, in *upload.ConfirmUploadRequest) (*upload.ConfirmUploadResponse, error) {
	fileName:=in.GetFileName()
	fmt.Printf("Master Confirmed File %s being uploaded successfully\n",fileName)
	return &upload.ConfirmUploadResponse{},nil
}

func handleUploadFile(path string,socket string){

	// Check if the file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist\n", path)
		return
	}

	filename := filepath.Base(path)

	//1. Establish Connection to the Master
	masterAddress := utils.GetMasterIP("client")
	connToMaster, err := grpc.Dial(masterAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Can not connect to Master at %s\n", masterAddress)
		return
	}
	fmt.Printf("Connected To Master %s\n", masterAddress)

	//2. Register as Client to Service Upload File offered by the Master
	uploadClient := upload.NewUploadServiceClient(connToMaster)


	// 3. Upload File Request
	fmt.Print("Sending Upload Request To Master ....\n")
	response, err:=uploadClient.RequestUpload(context.Background(),&upload.RequestUploadRequest{
		FileName: filename,
		ClientSocket: socket, 
	})
	if err!=nil{
		fmt.Println("Failed to request port from Master", err)
		return
	}

	nodeSocket:=response.GetNodeSocket()
	//Close Connection with Master
	connToMaster.Close()
	fmt.Printf("Sending File to %s ...\n", nodeSocket)


	//4. Transfer File
	//Establish Connection to Data Node
	connToDataNode, err := grpc.Dial(nodeSocket, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Could not Connect to DataNode at [%s]\n", nodeSocket)
		return
	}
	defer connToDataNode.Close()

	// Register To File Transfer Service to Data Node
	uploadClient=upload.NewUploadServiceClient(connToDataNode)

	// Send File MetaData + Chunks
	sendFile(path,uploadClient)
}

func sendFile(path string,uploadClient upload.UploadServiceClient){
	fileName := filepath.Base(path)

	//1. Open File
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Cannot open Video File at [%s] got error: %v\b", path ,err)
		return
	}
	defer file.Close()

	//2. Calling RPC
	stream, err := uploadClient.UploadFile(context.Background())
	if err != nil {
		fmt.Println("Cannot upload Video File: ", err)
		return
	}

	//3. Send MetaData For Video File
	req:=&upload.UploadFileRequest{
		Data: &upload.UploadFileRequest_FileInfo{
			FileInfo: &upload.FileInfo{
				FileName: fileName,
			},
		},
	}
	err = stream.Send(req)
	if err != nil {
		fmt.Println("cannot send file info to server: ", err, stream.RecvMsg(nil))
		return
	}

	//4. Send File Data
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			fmt.Println("End of File")
			// End Of File
			break
		}
		if err != nil {
			fmt.Println("cannot read chunk to buffer: ", err)
			return
		}
		// fmt.Printf("chunk of size %d n:%d\n", len(buffer[:n]), n)

		// Send New Request to Server with Data Chunk
		req:=&upload.UploadFileRequest{
			Data: &upload.UploadFileRequest_ChuckData{
				ChuckData: buffer[:n],
			},
		}
		err = stream.Send(req)
		if err != nil {
			fmt.Println("cannot send chunk to server: ", err, stream.RecvMsg(nil))
			return
		}
	}
	// [FIX] Same Port Used by Client to call the DataNode
	// fmt.Println("LOPPPS")
	// for{}
	//Sending EndOfFile
	stream.CloseAndRecv()

	fmt.Println("Finished Sending File 🧨")
}


func main() {
	fmt.Println("Welcome Client 😊")

	// [TODO] 
	// 1. Get Ip & Ports
	// ip,ports:=GetNodeSockets()
	ip:="127.0.0.1"
	port:="8085"
	client_socket:=ip+":"+port

	// client:=newClientNode(client_socket)

	// Create a channel for synchronization
	// notifyChan := make(chan struct{})

	for {
		// Transfer type from user
		// choose 1 for upload and 2 for download
		fmt.Print("Enter 1 for upload and 2 for download: ")
		var choice int
		fmt.Scanln(&choice)

		if choice != 1 && choice != 2 {
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter file path: ")
			var path string
			fmt.Scanln(&path)
			fmt.Print("Sending Upload Request ....\n")
			// (1) File Transfer Request
			// Send MetaData For Video File
			response, err := file_request_transfer_client.UploadRequest(context.Background(), &req.UploadFileRequest{
				FileName: filepath.Base(path),
			})
			if err != nil {
				fmt.Println("cannot request port from master to send the file", err)
				break
			}
			dataNodeAddress := response.Address

			// (2) File Transfer
			handleUploadFile(dataNodeAddress, path)

			//(3) Await Master Confirmation
			fmt.Println("Waiting For Confirm From Master")
			_, err = confirm_file_transfer_client.ConfirmFileTransfer(context.Background(), &cf.ConfirmFileTransferRequest{
				FileName: filepath.Base(path),
			})
			if err != nil {
				fmt.Println("Upload Failed please Try again")
				fmt.Println(err)
			} else {
				fmt.Println("Video Uploaded Successfully 🎆")
			}
			//Upload File
			handleUploadFile(path,client_socket)

			// // Confirm File Upload
			// master_listener, err := net.Listen("tcp", client_socket)
			// if err != nil {
			// 	fmt.Printf("Failed to Listen to %s client_socket\n",client_socket)
				
			// }
			// fmt.Printf("Listening to Master at Socket: %s\n",client_socket)
			
			
			// s := grpc.NewServer()
			// // client := &clientNode{}
			
			// // Register to UploadService
			// upload.RegisterUploadServiceServer(s, &client)
			
			// if err := s.Serve(master_listener); err != nil {
			// 	fmt.Println(err)
			// }
			// // [TODO] Add Channel
			// for {}
			// master_listener.Close()
			

			// ...
		
			// //(3) Await Master Confirmation
			// fmt.Println("Waiting For Confirm From Master")
			// _, err = confirm_file_transfer_client.ConfirmFileTransfer(context.Background(), &cf.ConfirmFileTransferRequest{
			// 	FileName: filepath.Base(path),
			// })
			// if err != nil {
			// 	fmt.Println("Upload Failed please Try again")
			// 	fmt.Println(err)
			// } else {
			// 	fmt.Println("Video Uploaded Successfully 🎆")
			// }

		case 2:
			fmt.Print("Enter file name: ")
			var filename string
			fmt.Scanln(&filename)
			fmt.Print("Downloading ....\n")
			// (1) check if file exists on master
			response, err := download_request_transfer_client.GetServer(context.Background(), &download.DownloadRequest{
				FileName: filename,
			})
			if err != nil {
				fmt.Println("cannot request port from master to download the file", err)
				break
			}
			// call get data from response
			file_exists := response.GetError()
			if file_exists != "" {
				fmt.Println("File does not exist")
				break
			}
			// (2) File Transfer
			servers := response.GetServers()
			if servers == nil {
				fmt.Println("No servers available")
				break
			}
			
			servers_list := servers.Servers
			// print servers list
			fmt.Println("servers list",servers_list)
		}
	}
}
