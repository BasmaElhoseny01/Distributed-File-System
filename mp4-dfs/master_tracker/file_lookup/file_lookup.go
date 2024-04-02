package file_lookup

import (
	"fmt"
	"sync"
	"time"
)

// Enum for the file Status
type FileStatus int
const (
    None FileStatus = iota
    Confirming
    Replicating
    // Add more status values as needed
)
func (fs FileStatus) String() string {
    switch fs {
    case None:
        return "None"
    case Confirming:
        return "Confirming"
    case Replicating:
        return "Replicating"
    default:
        return "Unknown"
    }
}

type File struct {
	// Define fields as needed
	file_name      string
	data_node_1    string
	replica_node_2 string
	replica_node_3 string

	confirmed bool
	replicating_nodes map[string]time.Time // Map with string keys and time.Time values
	
	// REMOVE
	status  FileStatus
	status_last_modified_timestamp time.Time
}

func NewFile(file_name string, data_node_1 string,path_1 string) File {
	return File{
		file_name:    file_name,

		data_node_1:  data_node_1,
		replica_node_2: "",
		replica_node_3: "",

		confirmed:false,
		replicating_nodes: make(map[string]time.Time), // Initialize the map

		// REMOVE
		status:None,
		status_last_modified_timestamp:time.Now(),
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
func (store *FileLookUpTable)ReplicateFile(file_name string ,node_id string)(error){
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if(store.data[file_name].data_node_1==""){
		store.data[file_name].data_node_1=node_id
	}else if(store.data[file_name].replica_node_2==""){
		store.data[file_name].replica_node_2=node_id
	}else if(store.data[file_name].replica_node_3==""){
		store.data[file_name].replica_node_3=node_id
	}
	 return nil
	
}

//Get Info
func (store *FileLookUpTable) PrintFileInfo(fileName string )(string){
	file:=store.data[fileName]

	details := fmt.Sprintf("[File] Name: %s,confirmed : %t,at node [%s] ,at node [%s] ,at node [%s], Status:[%s], Status Last Modified: [%s]",
	file.file_name,file.confirmed, file.data_node_1,file.replica_node_2,file.replica_node_3,file.status.String(),file.status_last_modified_timestamp.Format(time.RFC3339))
	return details
}

//Confirm File
func(store *FileLookUpTable) ConfirmFile(fileName string)(){
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.data[fileName].confirmed=true
	return
}


//Set File Operation Sate to None
func(store *FileLookUpTable) SetFileStateToNone(fileName string)(){
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.data[fileName].status=None

	// Update Timestamp for last status modification
	store.data[fileName].status_last_modified_timestamp=time.Now()

	return
}

//Set File Operation Sate to Confirming
func(store *FileLookUpTable) SetFileStateToConfirming(fileName string)(){
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.data[fileName].status=Confirming

	// Update Timestamp for last status modification
	store.data[fileName].status_last_modified_timestamp=time.Now()

	return
}

//Set File Operation Sate to Replicating
func(store *FileLookUpTable) SetFileStateToReplicating(fileName string)(){
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.data[fileName].status=Replicating

	// Update Timestamp for last status modification
	store.data[fileName].status_last_modified_timestamp=time.Now()

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
		if !file.confirmed && file.status!=Confirming {
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
		// [TODO] Fix IDLE Nodes check id node is idle
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


//Get Source Machines
func(store *FileLookUpTable) GetFileSourceMachines(fileName string)([]string){
	// get all possible Source Machine for the file
	store.mutex.Lock()
	defer store.mutex.Unlock()

	sourceMachines := make([]string, 0)

	file:=store.data[fileName]
	//TODO:fix IDILE

	if file.data_node_1 !=""{
		sourceMachines=append(sourceMachines,file.data_node_1)
	}

	if file.replica_node_2 !=""{
		sourceMachines=append(sourceMachines,file.replica_node_2)
	}

	if file.replica_node_3 !=""{
		sourceMachines=append(sourceMachines,file.replica_node_3)
	}
	return sourceMachines
}


//Get IdleFiles Files
func(store *FileLookUpTable) ResetFilesIdleStatus()([] string){
	store.mutex.Lock()
	defer store.mutex.Unlock()

	reset_files:= make([]string, 0)


	// for _, file := range store.data {
	// 	// [TODO] Fix IDLE Nodes
	// 	if file.status==Replicating || file.status==Confirming {
	// 		if  time.Since(file.status_last_modified_timestamp).Seconds()>20{
	// 			// Set Back to None
	// 			store.SetFileStateToNone(file.file_name)

	// 			reset_files=append(reset_files, file.file_name)
	// 		}
	// 	}
	// }
	return reset_files
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