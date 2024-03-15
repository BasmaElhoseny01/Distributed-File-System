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
	// replica_node_2 string
	// path_2         string
	// replica_node_3 string
	// path_3         string
}

func NewFile(file_name string, data_node_1 string,path_1 string) File {
	return File{
		file_name:    file_name,
		data_node_1:  data_node_1,
		path_1:       path_1,
		// replica_node_2: "-1",
		// path_2:"",
		// replica_node_3: "-1",
		// path_3:"",
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