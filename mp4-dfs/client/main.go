package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	cf "mp4-dfs/schema/confirm_file_transfer"
	tr "mp4-dfs/schema/file_transfer"
	req "mp4-dfs/schema/file_transfer_request"
	download "mp4-dfs/schema/download"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
)

func handleUploadFile(dataNodeAddress string, path string) {
	//Establish Connection to Data Node
	connToDataNode, err := grpc.Dial(dataNodeAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Could not Connect to Master at [%s]", dataNodeAddress)
		fmt.Println(err)
		return
	}
	// defer connToDataNode.Close()
	fmt.Printf("Connected To Data Node at %s 🤗\n", dataNodeAddress)

	// Register To File Transfer Service to Data Node
	file_transfer_client := tr.NewFileTransferServiceClient(connToDataNode)

	// File Transfer
	fmt.Printf("Sending File to Data Node ....\n")

	// Check if the file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist\n", path)
		return
	}

	// Open File
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Cannot open Video File at [%s] ", path)
		fmt.Println(err)
		return
	}
	defer file.Close()

	stream, err := file_transfer_client.UploadFile(context.Background())
	if err != nil {
		fmt.Println("Cannot upload Video File: ", err)
		return
	}

	// Send MetaData For Video File
	fileName := filepath.Base(path)

	req := &tr.UploadVideoRequest{
		Data: &tr.UploadVideoRequest_Info{
			Info: &tr.VideoInfo{
				Name: fileName,
			},
		},
	}
	err = stream.Send(req)
	if err != nil {
		fmt.Println("cannot send file info to server: ", err, stream.RecvMsg(nil))
		return
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)
	i := 0
	for {
		i += 1
		fmt.Println("Reading next buffer")
		n, err := reader.Read(buffer)
		if err == io.EOF {
			fmt.Printf("End of File\n\n")
			// End Of File
			break
		}
		if err != nil {
			fmt.Println("cannot read chunk to buffer: ", err)
			return
		}
		fmt.Printf("chunk of size %d n:%d\n", len(buffer[:n]), n)
		// Send New Request to Server with Data Chunk
		req := &tr.UploadVideoRequest{
			Data: &tr.UploadVideoRequest_ChuckData{
				ChuckData: buffer[:n],
			},
		}
		err = stream.Send(req)
		if err != nil {
			fmt.Println("cannot send chunk to server: ", err, stream.RecvMsg(nil))
			return
		}
	}

	// sends an EOF (End-of-File) signal to the server
	// does not directly close the server-side stream
	_, err = stream.CloseAndRecv()
	if err != nil {
		fmt.Print("cannot receive response:\n", err)
	}
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

	// File Transfer Confirm Request Service to Master
	confirm_file_transfer_client := cf.NewConfirmFileTransferServiceClient(connToMaster)

	download_request_transfer_client := download.NewDownloadServiceClient(connToMaster)



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
