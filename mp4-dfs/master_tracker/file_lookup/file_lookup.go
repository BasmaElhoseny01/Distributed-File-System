package file_lookup

import (
	"fmt"
	"sync"
)

type File struct {
	// Define fields as needed
	file_name      string
	data_node_1    string
	path_1         string
	replica_node_2 string
	path_2         string
	replica_node_3 string
	path_3         string

	confirmed bool
}

func NewFile(file_name string, data_node_1 string,path_1 string) File {
	return File{
		file_name:    file_name,
		data_node_1:  data_node_1,
		path_1:       path_1,
		replica_node_2: "-1",  //[FIX]
		path_2:"", //[FIX]
		replica_node_3: "-1", //[FIX]
		path_3:"", //[FIX],
		confirmed:false,
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

func (store *FileLookUpTable)GetFile(file_name string) (bool,string,string,string,string,string,string){
	// check if file already exist
	_, exists := store.data[file_name]
	if !exists {
		return false,"","","","","",""
	}
	return true,store.data[file_name].data_node_1,store.data[file_name].path_1,store.data[file_name].replica_node_2,store.data[file_name].path_2,store.data[file_name].replica_node_3,store.data[file_name].path_3
	
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

	details := fmt.Sprintf("[File] Name: %s,confirmed : %t,at node [%s] in %s ,at node [%s] in %s ,at node [%s] in %s",
	file.file_name,file.confirmed, file.data_node_1, file.path_1,file.replica_node_2,file.path_2,file.replica_node_3,file.path_3)
	return details
}

//GetNonConfirmed Files
func(store *FileLookUpTable) CheckUnConfirmedFiles()([]string){
	store.mutex.Lock()
	defer store.mutex.Unlock()

	nonConfirmedFiles := make([]string, 0)

	for _, file := range store.data {
		if !file.confirmed {
			nonConfirmedFiles = append(nonConfirmedFiles, file.file_name)
		}
	}

	return nonConfirmedFiles
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
