package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"
	"time"

	// "net"
	"os"
	"path/filepath"

	"google.golang.org/grpc"

	download "mp4-dfs/schema/download"
	upload "mp4-dfs/schema/upload"
	utils "mp4-dfs/utils"
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


var confirmationReceived bool // Shared variable to track confirmation status
var fileReceived string // Shared variable to track confirmation status
var confirmationMutex sync.Mutex // Mutex for concurrent access to the confirmation status

// ############################################# Rpcs #################################################################
// ConfirmUpload rpc
func (c *clientNode) ConfirmUpload(ctx context.Context, in *upload.ConfirmUploadRequest) (*upload.ConfirmUploadResponse, error) {
	// if break_client{
	// 	fmt.Println("Shutting Down ...")
	// 	os.Exit(0)
	// }

	confirmationMutex.Lock()
	defer confirmationMutex.Unlock()
	if fileReceived == ""{
		return &upload.ConfirmUploadResponse{Status: "time_out"},nil
	}

	confirmationReceived = true
	fileName:=in.GetFileName()
	if fileName!=fileReceived{
		return &upload.ConfirmUploadResponse{Status: "wrong_file"},nil
	}
	
	fmt.Printf("Master Confirmed File %s being uploaded successfully\n",fileName)
	return &upload.ConfirmUploadResponse{Status: "success"},nil
}


//  ######################################################## GO Routines ######################################################
func handleUploadFile(path string,socket string) error{

	// Check if the file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist\n", path)
		return err
	}

	filename := filepath.Base(path)

	//1. Establish Connection to the Master
	masterAddress := utils.GetMasterIP("client")
	connToMaster, err := grpc.Dial(masterAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Can not connect to Master at %s\n", masterAddress)
		return err
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
		return err
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
		return err
	}
	defer connToDataNode.Close()

	// Register To File Transfer Service to Data Node
	uploadClient=upload.NewUploadServiceClient(connToDataNode)

	// Send File MetaData + Chunks
	sendFile(path,uploadClient)
	return nil
}


func handleDownloadFile(servers_data []*download.Server, filename string){
	// 1. Establish Connection to data node
	if len(servers_data) == 1 {
		fmt.Println("Only 1 server available")
		// connect to the server
		data_node := servers_data[0]

		address := data_node.Ip + ":" + data_node.Port

		// filepath := data_node.FilePath BASMA

		connToDataNode, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			fmt.Println("Cannot connect to Data Node at", address)
			return
		}
		defer connToDataNode.Close()

		// send download request
		downloadClient := download.NewDownloadServiceClient(connToDataNode)
		stream, err := downloadClient.Download(context.Background(), &download.DownloadRequest{
			FileName: filename,
		})
		if err != nil {
			fmt.Println("Cannot download file", err)
			return
		}

		// create file
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("Cannot create file", err)
			return
		}

		// receive chunks
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
		fmt.Println("File downloaded successfully")
	} else {
		fmt.Println("Multiple servers available")
	}
}

// ############################################################ Utils ############################################################
// Send File
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
	fmt.Println("Sending End Of File To DataNode")
	_,err=stream.CloseAndRecv()
	if err !=nil{
		fmt.Println("Can not receive Upload File Response from DataNode ", err, stream.RecvMsg(nil))
		return
	}
	fmt.Println("Finished Sending File 🧨")
}

func GetNodeSockets() (node_ip string, port string,breakEnabled bool) {
	// Define flags
	var breakFlag bool
	flag.BoolVar(&breakFlag, "break", false, "Enable break")

	// Parse flags
	flag.Parse()

	if breakFlag{
		// Remove the element at index 2
		os.Args = append(os.Args[:1], os.Args[2:]...)
	}

	if(len(os.Args)<=1){
		// go run ./client/main.go --> MyIP + Empty Port [Done]

		// Get IP address using GetMyIP function if not provided in arguments
		ip := utils.GetMyIp()
		if ip == nil {
			fmt.Println("Failed to retrieve IP address.")
			os.Exit(1)
		}
		//No ports are passed then get empty one
		port_no,err:=utils.GetEmptyPort(ip.String(),"-1")
		if err!=nil{
			fmt.Printf("Error When Getting Empty Port for the Client Error:%v",err)
			os.Exit(1)
		}

		return ip.String(),port_no,breakFlag
	}


	if len(os.Args) > 4 {
        fmt.Println("Usage: client [-break] [<your_ip>] [<port>]")
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
	// Parse Port if passed or get Empty One
	if len(os.Args)>1{
		// Then a port is passed
		// Parse Port
		p_in, err := strconv.Atoi(os.Args[1])
		fmt.Println("port",p_in)
		if err != nil || p_in <= 0 || p_in > 65535 {
			fmt.Println("Invalid Port:", os.Args[len(os.Args)-1])
			os.Exit(1)
		}

		// Check if Port is reachable
		if !utils.IsPortOpen(ip.String(), p_in) {
			fmt.Printf("Port %d is not reachable\n", p_in)
			os.Exit(1)
		}
		port=strconv.Itoa(p_in)
	}else{
		//No ports are passed then get empty one
		p,err:=utils.GetEmptyPort(ip.String(),"-1")
		if err!=nil{
			fmt.Printf("Error When Getting Empty Port for the Client Error:%v",err)
			os.Exit(1)
		}
		port=p
	}


	return ip.String(),port,breakFlag
}

// var break_client bool

func main() {
	fmt.Println("Welcome Client 😊")
	// [TODO] 
	// 1. Get Ip & Ports
	// [Fix] this Func to be like that in DataNode
	ip,port,breakEnabled:=GetNodeSockets()
	// ip:="127.0.0.1"
	// port := "12874"
	client_socket:=ip+":"+port

	client:=newClientNode(client_socket)

	master_listener, err := net.Listen("tcp", client_socket)
	if err != nil {
		fmt.Printf("Failed to Listen to %s\n",client_socket)
		
	}
	defer master_listener.Close()
	fmt.Printf("Listening to Socket: %s\n",client_socket)
	if breakEnabled{
		// break_client=true
		fmt.Println("You asked me to Break :(",breakEnabled)
	}


	
	s := grpc.NewServer()
	// client := &clientNode{}

	// Register to UploadService
	upload.RegisterUploadServiceServer(s, &client)

	//Serve Calls in a go routine
	go func ()  {
		if err := s.Serve(master_listener); err != nil {
			fmt.Println(err)
		}
	}()
	
	for {
		// Transfer type from user
		// choose 1 for upload and 2 for download
		fmt.Print("Enter 1 for upload and 2 for download:")
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
			fileReceived=filepath.Base(path)

			// Add a variable to track whether confirmation is received
			confirmationMutex.Lock()
			confirmationReceived = false
			confirmationMutex.Unlock() 

			//Upload File
			err:=handleUploadFile(path,client_socket)
			if err!=nil{
				confirmationMutex.Lock()
				fileReceived=""
				confirmationMutex.Unlock()
				continue
			}

			//Wait For Confirm From Master
			fmt.Println("Waiting for confirmation from server...")
			attempts_count:=0
			for{
				confirmationMutex.Lock()

				attempts_count+=1
				if confirmationReceived {
					confirmationMutex.Unlock()
					fmt.Println("Video Uploaded Successfully 🎆")
					break
				}
				if attempts_count>5{
				// if attempts_count>2{
					fileReceived=""
					confirmationMutex.Unlock()
					fmt.Println("Upload Failed please Try again")
					break
				}

				confirmationMutex.Unlock()
				// Sleep for 5 seconds before the next check
				time.Sleep(5 * time.Second)
			}
		
		case 2:
			fmt.Print("Enter file name: ")
			var filename string
			fmt.Scanln(&filename)
			fmt.Print("Downloading ....\n")

			//1. Establish Connection to the Master
			masterAddress := utils.GetMasterIP("client")
			connToMaster, err := grpc.Dial(masterAddress, grpc.WithInsecure())
			if err != nil {
				fmt.Println(err)
				fmt.Printf("Can not connect to Master at %s\n", masterAddress)
				return
			}
			fmt.Printf("Connected To Master %s\n", masterAddress)
			download_request_transfer_client := download.NewDownloadServiceClient(connToMaster)

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
			// close connection to master
			connToMaster.Close()

			// (3) Download File
			handleDownloadFile(servers_list, filename)
		}
	}
}
// go run ./client/main.go --> MyIP + Empty Port [Done]
// go run ./client/main.go 127.0.0.1 --> 127.0.0.1 + Empty Port [Done]
// go run ./client/main.go 127.0.0.1 8090 --> 127.0.0.1 + 8090 [Done]
// go run ./client/main.go 8090 --> MyIP + 8090 [Done]


// go run ./client/main.go -break --> MyIP + Empty Port [Done]
// go run ./client/main.go -break 127.0.0.1 --> 127.0.0.1 + Empty Port [Done]
// go run ./client/main.go -break 127.0.0.1 8090 --> 127.0.0.1 + 8090 [Done]
// go run ./client/main.go -break 8090 --> MyIP + 8090 [Done]
