package file_lookup

import (
	"fmt"
	"sync"
	"time"
)

type File struct {
	// Define fields as needed
	file_name      string
	data_node_1    string
	replica_node_2 string
	replica_node_3 string

	confirmed bool
	
	// State Management
	replicating_nodes map[string]time.Time // Map with string keys and time.Time values
	confirming bool
	confirming_timestamp time.Time
}

func NewFile(file_name string, data_node_1 string) File {
	return File{
		file_name:    file_name,

		data_node_1:  data_node_1,
		replica_node_2: "",
		replica_node_3: "",

		confirmed:false,

		replicating_nodes: make(map[string]time.Time), // Initialize the map
		confirming: false,
		confirming_timestamp: time.Now(),

	}
}

type FileLookUpTable struct {
	mutex sync.RWMutex
	data  map[string]*File //Map of key are strings and values are Files
}

func NewFileLookUpTable() FileLookUpTable {
	return FileLookUpTable{
		data: make(map[string]*File),
	}
}

func (store *FileLookUpTable)CheckFileExists(file_name string) (bool){
	_, exists := store.data[file_name]
    return exists
}

func (store *FileLookUpTable)GetFile(file_name string) (bool,string,string,string){
	// check if file already exist
	_, exists := store.data[file_name]
	if !exists {
		return false,"","",""
	}
	return true,store.data[file_name].data_node_1,store.data[file_name].replica_node_2,store.data[file_name].replica_node_3
	
	// return true,store.data[file_name].data_node_1,store.data[file_name].path_1,store.data[file_name].replica_node_2,store.data[file_name].path_2,store.data[file_name].replica_node_3,store.data[file_name].path_3
}

//Add New File
func (store *FileLookUpTable)AddFile(mp4file *File) (error){
	// Add File to the lookup table
	store.mutex.Lock()
	defer store.mutex.Unlock()

	store.data[mp4file.file_name]=mp4file
	for key, value := range store.data {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}
	return nil
}
func (store *FileLookUpTable)ReplicateFile(file_name string ,node_id string,IsNodeAlive func(string) (bool))(error){
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if(store.data[file_name].data_node_1=="" || !IsNodeAlive(store.data[file_name].data_node_1)){
		store.data[file_name].data_node_1=node_id
	}else if(store.data[file_name].replica_node_2=="" || !IsNodeAlive(store.data[file_name].replica_node_2)){
		store.data[file_name].replica_node_2=node_id
	}else if(store.data[file_name].replica_node_3=="" || !IsNodeAlive(store.data[file_name].replica_node_3)){
		store.data[file_name].replica_node_3=node_id
	}
	 return nil
	
}

//Get Info
func (store *FileLookUpTable) PrintFileInfo(fileName string )(string){
	file:=store.data[fileName]

	details := fmt.Sprintf("[File] Name: %s,confirmed : %t,at node [%s] ,at node [%s] ,at node [%s]",
	file.file_name,file.confirmed, file.data_node_1,file.replica_node_2,file.replica_node_3)
	return details
}

//Confirm File
func(store *FileLookUpTable) ConfirmFile(fileName string)(){
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.data[fileName].confirmed=true
	return
}


//Remove File
func(store *FileLookUpTable) RemoveFile(fileName string)(){
	store.mutex.Lock()
    defer store.mutex.Unlock()
    delete(store.data, fileName)
	return
}


//GetNonConfirmed Files
func(store *FileLookUpTable) CheckUnConfirmedFiles()([]string){
	store.mutex.Lock()
	defer store.mutex.Unlock()

	nonConfirmedFiles := make([]string, 0)

	for _, file := range store.data {
		// File is unconfirmed and not being confirmed now :D
		if !file.confirmed && !file.confirming {
			nonConfirmedFiles = append(nonConfirmedFiles, file.file_name)
		}
	}

	return nonConfirmedFiles
}

//GetUnReplicated Files
func(store *FileLookUpTable) CheckUnReplicatedFiles(IsNodeAlive func(string) (bool))([]string){
	store.mutex.Lock()
	defer store.mutex.Unlock()

	nonReplicatedFiles := make([]string, 0)

	count_non_replica:=0
	
	for _, file := range store.data {
		if (file.data_node_1 =="" || !IsNodeAlive(file.data_node_1)){
			count_non_replica+=1
		}
		if (file.replica_node_2 ==""||!IsNodeAlive(file.replica_node_2)){
			count_non_replica+=1
		}
		if (file.replica_node_3==""||!IsNodeAlive(file.replica_node_3)) {
			count_non_replica+=1
		}

		// Check if already Replicating
		if count_non_replica> len(store.data[file.file_name].replicating_nodes){
			nonReplicatedFiles = append(nonReplicatedFiles, file.file_name)
		}
		
	}
	return nonReplicatedFiles
}


//Get Replication Source Machines
func(store *FileLookUpTable) GetFileSourceMachines(fileName string,IsNodeAlive func(string) (bool))([]string){
	// get all possible Source Machine for the file
	store.mutex.Lock()
	defer store.mutex.Unlock()

	sourceMachines := make([]string, 0)

	file:=store.data[fileName]

	// Checking Non replicated Nodes + Idle Nodes
	if (file.data_node_1 !="" && IsNodeAlive(file.data_node_1)){
		sourceMachines=append(sourceMachines,file.data_node_1)
	}

	if (file.replica_node_2 !="" && IsNodeAlive(file.replica_node_2)){
		sourceMachines=append(sourceMachines,file.replica_node_2)
	}

	if (file.replica_node_3 !=""&& IsNodeAlive(file.replica_node_3)){
		sourceMachines=append(sourceMachines,file.replica_node_3)
	}
	return sourceMachines
}

func (store *FileLookUpTable) SetConfirming(file_name string) (error){
	store.mutex.Lock()
	defer store.mutex.Unlock()
	
	// Add the node_id to the replicating_nodes map with a default timestamp
    store.data[file_name].confirming = true
	store.data[file_name].confirming_timestamp = time.Now()
	return nil
}

func (store *FileLookUpTable) UnSetConfirming(file_name string) (error){
	store.mutex.Lock()
	defer store.mutex.Unlock()
	
	// Add the node_id to the replicating_nodes map with a default timestamp
    store.data[file_name].confirming = false
	return nil
}

func (store *FileLookUpTable) AddReplicatingNode(file_name string,node_id string) (error){
	// Add the node_id to the replicating_nodes map with a default timestamp
    store.data[file_name].replicating_nodes[node_id] = time.Now()
	return nil
}

func (store *FileLookUpTable) RemoveReplicatingNode(file_name string,node_id string){
	// Iterate over the map to find and remove the element with the specified node_id
	for key := range store.data[file_name].replicating_nodes {
		if key == node_id {
			// Remove the element with the matching node_id from the replicating_nodes map
			delete(store.data[file_name].replicating_nodes, key)

			// Assuming node_id is unique, break the loop after the removal
			break
		}
	}
}


func (store *FileLookUpTable) ResetIdleFiles(max float64){
	// Iterate over the map to find and Idle Files
	for _, file := range store.data {
		
		// Idle While Confirming
		if file.confirming && time.Since(file.confirming_timestamp).Seconds()>max{
			// Break Confirming
			store.UnSetConfirming(file.file_name)
			fmt.Printf("[Break Confirmation] File %s\n",file.file_name)
		}

		// Idle While Replicating
		for key := range file.replicating_nodes {
			last_time_stamp:=file.replicating_nodes[key]

			if time.Since(last_time_stamp).Seconds()>max{
				// Break Replicating 
				delete(file.replicating_nodes, key)
				fmt.Printf("[Break Replication] File %s By Node %s\n",file.file_name,key)
			}
		}
	}
}