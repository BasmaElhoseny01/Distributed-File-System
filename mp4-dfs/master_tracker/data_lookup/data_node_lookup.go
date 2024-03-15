package data_lookup

import (
	"strconv"
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
	port string
}

func NewDataNode(Ip string,port string) DataNode{
	return DataNode{
		Id: "",
		Ip: Ip,
		port: port,
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
	address:=store.data[min_node].Ip+":"+store.data[min_node].port
	
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
	port:=store.data[id].port
	
	return Ip,port
}