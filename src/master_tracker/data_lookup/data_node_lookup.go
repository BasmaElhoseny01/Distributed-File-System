package data_lookup

import (
	"fmt"
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
	file_port string
	replication_port string
}

func NewDataNode(Ip string,file_port string,replication_port string) DataNode{
	return DataNode{
		Id: "",
		Ip: Ip,
		file_port:file_port ,
		replication_port: replication_port,
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

	details := fmt.Sprintf("[DataNode] ID: %s, IP: %s, File Port: %s, Replica Port: %s, Alive: %v, Ping Timestamp: %s, Load: %.2f",
	data_node.Id, data_node.Ip, data_node.file_port,data_node.replication_port, data_node.alive, data_node.ping_timestamp.Format(time.RFC3339), data_node.load)
	return details
}

// Update DataNode Time Stamp
func (store *DataNodeLookUpTable)UpdateNodeTimeStamp(id string) (time.Time,error){
	store.mutex.Lock()
	defer store.mutex.Unlock()

	store.data[id].ping_timestamp=time.Now()
	
	return store.data[id].ping_timestamp,nil
}

// Update DataNode 
func (store *DataNodeLookUpTable)UpdateNodeLoad(id string) (error){
	store.mutex.Lock()
	defer store.mutex.Unlock()

	store.data[id].load+=1.0
	
	return nil
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
	// address:=store.data[min_node].Ip+":"+store.data[min_node].ports[0]
	
	return min_node,nil
}

// Update DataNode status
func (store *DataNodeLookUpTable)UpdateNodeStatus(id string,status bool) error{
	store.mutex.Lock()
	defer store.mutex.Unlock()
	
	store.data[id].alive=status
	
	return nil
}

// get ip and port for replication services of a node
func (store *DataNodeLookUpTable)GetNodeFileServiceAddress(id string) (string,string){
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	
	if id==""{
		return "",""
	}
	// check if node alive
	if !store.data[id].alive {
		return "",""
	}
	Ip:=store.data[id].Ip
	port:=store.data[id].file_port
	
	return Ip,port
}

// get ip and port for file services of a node
func (store *DataNodeLookUpTable)GetNodeReplicationServiceAddress(id string) (string,string,error){
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	
	// check if node alive
	if !store.data[id].alive {
		return "","",nil
	}
	Ip:=store.data[id].Ip
	port:=store.data[id].replication_port
	
	return Ip,port,nil
}

// // get ip and port of a node
// func (store *DataNodeLookUpTable)GetNodeAddress(id string,port_type string) (string,string){
// 	store.mutex.RLock()
// 	defer store.mutex.RUnlock()
	
// 	// check if node alive
// 	if !store.data[id].alive {
// 		return "",""
// 	}
// 	Ip:=store.data[id].Ip
// 	if port_type=="file"
// 	port:=store.data[id].ports[0]
// 	else if
	
// 	return Ip,port
// }

func (store *DataNodeLookUpTable) CheckPingStatus(){
	store.mutex.Lock()
	defer store.mutex.Unlock()

	for _, dataNode := range store.data {
		if dataNode.alive && time.Since(dataNode.ping_timestamp).Seconds()>4{
			fmt.Printf("DataNode %s is Idl at time stamp %s \n",dataNode.Id,dataNode.ping_timestamp.Format(time.RFC3339))
			dataNode.alive=false
		}
	}
}


//Replicas Functions
func (store *DataNodeLookUpTable)GetCopyDestination(sources []string) (string,error){
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	// Get Destinations that aren't in the source list

	for _, value := range store.data {
		if value.alive {
			exist:=false
			for _,source :=range sources{
				if(source == value.Id){
					exist=true
					break
				}
			}
			if !exist{
				return value.Id,nil
			}

		}
	}
	return "",nil
}


func (store *DataNodeLookUpTable)IsNodeAlive(node_id string) (bool){
	return store.data[node_id].alive
}