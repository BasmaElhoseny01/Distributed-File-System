package file_lookup

import (
	"fmt"
	"sync"
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
	status  FileStatus
}

func NewFile(file_name string, data_node_1 string,path_1 string) File {
	return File{
		file_name:    file_name,

		data_node_1:  data_node_1,
		replica_node_2: "-1",  //[FIX]
		replica_node_3: "-1", //[FIX]

		confirmed:false,
		status:None,
		// [TODO] Add last time state changed
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

//Get Info
func (store *FileLookUpTable) PrintFileInfo(fileName string )(string){
	file:=store.data[fileName]

	details := fmt.Sprintf("[File] Name: %s,confirmed : %t,at node [%s] ,at node [%s] ,at node [%s], Status:[%s]",
	file.file_name,file.confirmed, file.data_node_1,file.replica_node_2,file.replica_node_3,file.status.String())
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
	return
}

//Set File Operation Sate to Confirming
func(store *FileLookUpTable) SetFileStateToConfirming(fileName string)(){
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.data[fileName].status=Confirming
	return
}

//Set File Operation Sate to Replicating
func(store *FileLookUpTable) SetFileStateToReplicating(fileName string)(){
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.data[fileName].status=Replicating
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
func(store *FileLookUpTable) CheckUnReplicatedFiles()([]string){
	store.mutex.Lock()
	defer store.mutex.Unlock()

	nonReplicatedFiles := make([]string, 0)

	for _, file := range store.data {
		// [TODO] Fix IDLE Nodes
		if (file.replica_node_2 =="-1" ||  file.replica_node_3=="-1") && file.status!=Replicating {
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

	if file.data_node_1 !="-1"{
		sourceMachines=append(sourceMachines,file.data_node_1)
	}

	if file.replica_node_2 !="-1"{
		sourceMachines=append(sourceMachines,file.replica_node_2)
	}

	if file.replica_node_3 !="-1"{
		sourceMachines=append(sourceMachines,file.replica_node_3)
	}
	return sourceMachines
}