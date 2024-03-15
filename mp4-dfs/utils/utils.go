package utils

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	// "time"
)

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
        Port int    `json:"port"`
    } `json:"master"`
}


func GetMasterIP() string{
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

    return fmt.Sprintf("%s:%s", configInfo.Master.IP, strconv.Itoa(configInfo.Master.Port))
}