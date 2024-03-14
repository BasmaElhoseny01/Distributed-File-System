package main

import (
	"context"
	"fmt"

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
	// File Transfer Service to Data Node


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

			// TODO (2) File Transfer
			fmt.Printf("Sending File to [%s] ....\n",response)

		case 2:
			fmt.Print("Downloading ....\n")
		}

	}
}
