package main

import (
	"context"
	"fmt"

	tr "mp4-dfs/schema/file_transfer"
	req "mp4-dfs/schema/file_transfer_request"

	"google.golang.org/grpc"
)



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
		var filepath string
		fmt.Scanln(&filepath)
		fmt.Print(filepath)

		switch choice {
		case 1:
			fmt.Print("Sending Upload Request ....\n")
			// TODO (1) File Transfer Request
			response,err :=file_request_transfer_client.UploadRequest(context.Background(), &req.UploadFileRequest{})
			if err != nil {
				fmt.Println(err)
			}

			dataNodePort:=response.Address

			//Establish Connection to Data Node
			connToDataNode, err := grpc.Dial(dataNodePort, grpc.WithInsecure())
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Connected To Data Node at %s 🤗\n",dataNodePort)

			// Register To File Transfer Service to Data Node
			file_transfer_client := tr.NewFileTransferServiceClient(connToDataNode)

			// File Transfer
			fmt.Printf("Sending File to Data Node ....\n")

			stream,err:=file_transfer_client.UploadFile(context.Background())

			
			fmt.Print("\nstream",stream)
			fmt.Print("\nerr",err)

			// Disconnect with Data Node
			connToMaster.Close()

		case 2:
			fmt.Print("Downloading ....\n")
		}

	}
}
