package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	tr "mp4-dfs/schema/file_transfer"
	req "mp4-dfs/schema/file_transfer_request"

	"google.golang.org/grpc"
)


func handleUploadFile(dataNodePort string,path string){
	//Establish Connection to Data Node
	connToDataNode, err := grpc.Dial(dataNodePort, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	// defer connToDataNode.Close()
	fmt.Printf("Connected To Data Node at %s 🤗\n",dataNodePort)

	// Register To File Transfer Service to Data Node
	file_transfer_client := tr.NewFileTransferServiceClient(connToDataNode)

	// File Transfer
	fmt.Printf("Sending File to Data Node ....\n")

	// Check if the file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist\n", path)
	}

	// Open File
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Cannot open Video File: ", err)
	}
	fmt.Print(file)
	defer file.Close()
	
	stream,err:=file_transfer_client.UploadFile(context.Background())
	if err != nil {
		fmt.Println("Cannot upload Video File: ", err)
	}

	// Send MetaData For Video File
	fileName := filepath.Base(path)

	req:=&tr.UploadVideoRequest{
		Data: &tr.UploadVideoRequest_Info{
			Info: &tr.VideoInfo{
				Name: fileName,
			},
		},
	}
	err = stream.Send(req)
	if err != nil {
		fmt.Println("cannot send file info to server: ", err, stream.RecvMsg(nil))
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)
	i:=0
	for {
		i+=1;
		fmt.Println("Reading next buffer")
		n, err := reader.Read(buffer)
		if err == io.EOF {
			fmt.Printf("End of File\n\n")
			// End Of File
			break
		}
		if err != nil {
			fmt.Println("cannot read chunk to buffer: ", err)
		}
		fmt.Printf("chunk of size %d n:%d\n", len(buffer[:n]),n)
		// Send New Request to Server with Data Chunk
		req:= &tr.UploadVideoRequest{
			Data: &tr.UploadVideoRequest_ChuckData{
				ChuckData: buffer[:n],
			},
		}
		err = stream.Send(req)
		if err != nil {
			fmt.Println("cannot send chunk to server: ", err, stream.RecvMsg(nil))
		}

		// if i==2{
		// 	fmt.Println("BREEE")
		// 	break;
		// }
	}

	// sends an EOF (End-of-File) signal to the server
	// does not directly close the server-side stream
	_, err = stream.CloseAndRecv()
	if err != nil {
		fmt.Print("cannot receive response:\n", err)
	}

	//[TODO] Wait To Receive Confirm From Master
	fmt.Println("Waiting For Confirm From Master")
	for{}
	
	fmt.Println("Video Uploaded Successfully to Data Node")
}

func main() {
	fmt.Println("Welcome Client 😊")

	masterPort := "localhost:5001"
	connToMaster, err := grpc.Dial(masterPort, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer connToMaster.Close()
	fmt.Print("Connected To Master Node 🎉\n")


	// Register To Services
	// File Transfer Request Service to Master
	file_request_transfer_client := req.NewFileTransferRequestServiceClient(connToMaster)


	for{
		// Transfer type from user
		// choose 1 for upload and 2 for download
		fmt.Print("Enter 1 for upload and 2 for download: ")
		var choice int
		fmt.Scanln(&choice)

		if choice != 1 && choice != 2 {continue}


		fmt.Print("Enter file path: ")
		var path string
		fmt.Scanln(&path)
	

		switch choice {
		case 1:
			fmt.Print("Sending Upload Request ....\n")
			// TODO (1) File Transfer Request
			response,err :=file_request_transfer_client.UploadRequest(context.Background(), &req.UploadFileRequest{})
			if err != nil {
				fmt.Println(err)
			}

			dataNodePort:=response.Address
			handleUploadFile(dataNodePort,path)

		case 2:
			fmt.Print("Downloading ....\n")
		}

	}
}
