package data_lookup

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Define your struct for DataNode
type DataNode struct {
    // Define fields as needed
    Id  string

	alive bool 
	ping_timestamp time.Time
	load float32

	Ip string
	port string //[FIX] Remove
	ports []string
}

// [FIX] Remove port
func NewDataNode(Ip string,port string,ports []string) DataNode{
	return DataNode{
		Id: "",
		Ip: Ip,
		port: port,
		ports:ports,
	}
}

type DataNodeLookUpTable struct{
	mutex sync.RWMutex
	data map[string]*DataNode //Map of key are strings and values are DataNodes
	nodes_count int	

}

func NewDataNodeLookUpTable() DataNodeLookUpTable{
	return DataNodeLookUpTable{
		data:make(map[string]*DataNode),
	}
}

// Add New DataNode
func (store *DataNodeLookUpTable)AddDataNode(dataNode *DataNode) (string,error){
	store.mutex.Lock()
	defer store.mutex.Unlock()

	// Get Id for teh node
	store.nodes_count++
	dataNode.Id=strconv.Itoa(store.nodes_count) 

	// Set variables
	dataNode.alive=true
	dataNode.ping_timestamp=time.Now()
	dataNode.load=0.0

	// Add 	node to the lookup table
	store.data[dataNode.Id]=dataNode

	return dataNode.Id,nil
}

//Get Info
func (store *DataNodeLookUpTable) PrintDataNodeInfo(data_node_id string )(string){
	data_node:=store.data[data_node_id]

	details := fmt.Sprintf("DataNode ID: %s, IP: %s, Ports: %s, Alive: %v, Ping Timestamp: %s, Load: %.2f",
	data_node.Id, data_node.Ip, strings.Join(data_node.ports, "-"), data_node.alive, data_node.ping_timestamp.Format(time.RFC3339), data_node.load)
	return details
}

// Update DataNode Time Stamp
func (store *DataNodeLookUpTable)UpdateNodeTimeStamp(id string) (time.Time,error){
	store.mutex.Lock()
	defer store.mutex.Unlock()

	store.data[id].ping_timestamp=time.Now()
	
	return store.data[id].ping_timestamp,nil
}

// Get Least Loaded DataNode
func (store *DataNodeLookUpTable)GetLeastLoadedNode() (string,error){
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	min_load :=float32(-1.0)
	min_node:=""

	for key, value := range store.data {
		if (value.alive && value.load<min_load) || min_load==-1.0{
			min_load=value.load
			min_node=key
		}
	}

	// [FIX] Return the Available Port replace ports[0] with that 
	address:=store.data[min_node].Ip+":"+store.data[min_node].ports[0]
	
	return address,nil
}

// Update DataNode status
func (store *DataNodeLookUpTable)UpdateNodeStatus(id string,status bool) error{
	store.mutex.Lock()
	defer store.mutex.Unlock()
	
	store.data[id].alive=status
	
	return nil
}

// get ip and port of a node
func (store *DataNodeLookUpTable)GetNodeAddress(id string) (string,string){
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	
	// check if node alive
	if !store.data[id].alive {
		return "",""
	}
	Ip:=store.data[id].Ip
	port:=store.data[id].ports[0]
	
	return Ip,port
}