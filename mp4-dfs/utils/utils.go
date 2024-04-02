package utils

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	// "time"
)

// emptyFolder empties the contents of the given folder
func EmptyFolder(folderPath string) error {
    // Iterate through the files in the folder
    err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        // Remove regular files and directories (excluding the root folder itself)
        if info.Mode().IsRegular() || info.IsDir() && path != folderPath {
            return os.RemoveAll(path)
        }
        return nil
    })
    return err
}

func GetMyIp() (net.IP){
	interfaces, err := net.Interfaces()
    if err != nil {
        fmt.Println("Error:", err)
        return nil
    }

    var wifiInterface string
    for _, iface := range interfaces {
        if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 {
            addrs, err := iface.Addrs()
            if err != nil {
                fmt.Println("Error:", err)
                return nil
            }
            for _, addr := range addrs {
                if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.To4() != nil {
                    if iface.Flags&net.FlagBroadcast != 0 {
                        wifiInterface = iface.Name
                        break
                    }
                }
            }
        }
    }

    if wifiInterface == "" {
        fmt.Println("No active WiFi interface found.")
        return nil
    }

    iface, err := net.InterfaceByName(wifiInterface)
    if err != nil {
        fmt.Println("Error:", err)
        return nil
    }

    addrs, err := iface.Addrs()
    if err != nil {
        fmt.Println("Error:", err)
        return nil
    }

    var ip net.IP
    for _, addr := range addrs {
        ipnet, ok := addr.(*net.IPNet)
        if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
            ip = ipnet.IP
            break
        }
    }

	
    if ip != nil {
		return ip
    } else {
		return nil
    }
}

func IsPortOpen(ip string, port int) bool {
    address := fmt.Sprintf("%s:%d", ip, port)
    client_listener, err := net.Listen("tcp", address)
    
	if err != nil {
        fmt.Println("Error:",err)
        return false
	}
    defer client_listener.Close()
    return true
}

type ConfigInfo struct {
    Master struct {
        IP   string `json:"ip"`
        ClientPort int `json:"client_port"`
        NodePort int `json:"node_port"`
        NodePingPort int `json:"node_ping_port"`
    } `json:"master"`
}


func GetMasterIP(source string) (string){
    if source!="node" && source!="client" && source!="ping"{
        fmt.Printf("GetMasterIP Option must be node or source while %s was passed:\n", source)
        os.Exit(0)
    } 

    // Open the JSON file
    file, err := os.Open("config.json")
    if err != nil {
        fmt.Println("Error When Reading Config:", err)
        os.Exit(0)
    }
    defer file.Close()

    // Decode the JSON
    var configInfo ConfigInfo
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&configInfo)
    if err != nil {
        fmt.Println("Error When Reading Config:", err)
        os.Exit(0)
    }
    if source=="node"{
        return fmt.Sprintf("%s:%s", configInfo.Master.IP, strconv.Itoa(configInfo.Master.NodePort))
    } else if source=="ping"{
        return fmt.Sprintf("%s:%s", configInfo.Master.IP, strconv.Itoa(configInfo.Master.NodePingPort))
        
    }
    return fmt.Sprintf("%s:%s", configInfo.Master.IP, strconv.Itoa(configInfo.Master.ClientPort))
}

func GetEmptyPort(ip string) (port string, err error) {
	for p := 1024; p <= 65535; p++ {
		addr := fmt.Sprintf("%s:%d", ip, p)
		// Attempt to listen on the port
		listener, err := net.Listen("tcp", addr)
		if err == nil {
			// Port is available, close the listener and return the port
			_ = listener.Close()
			return fmt.Sprintf("%d", p), nil
		}
	}
	return "", fmt.Errorf("no available ports found on %s", ip)
}