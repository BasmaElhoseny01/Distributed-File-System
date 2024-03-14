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
    // Add more fields as needed
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

	// Add 	node to the lookup table
	store.data[dataNode.Id]=dataNode

	return dataNode.Id,nil
}


// Update DataNode status
func (store *DataNodeLookUpTable)UpdateNodeTimeStamp(id string) (time.Time,error){
	store.mutex.Lock()
	defer store.mutex.Unlock()

	store.data[id].ping_timestamp=time.Now()
	
	return store.data[id].ping_timestamp,nil
}
// Update DataNode status
func (store *DataNodeLookUpTable)UpdateNodeStatus(id string,status bool) error{
	store.mutex.Lock()
	defer store.mutex.Unlock()
	
	store.data[id].alive=status
	
	return nil
}