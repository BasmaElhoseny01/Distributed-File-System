package file_lookup

import (
	"sync"
)

type File struct {
	// Define fields as needed
	file_name      string
	data_node_1    string
	replica_node_1 string
	replica_node_2 string
}

func NewFile(file_name string, data_node_1 string, replica_node_1 string, replica_node_2 string) File {
	return File{
		file_name:      file_name,
		data_node_1:    data_node_1,
		replica_node_1: "-1",
		replica_node_2: "-1",
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

// //Add New File
// func (store *FileLookUpTable)AddFile(
// 	mp4file *File,
// 	mp4Data bytes.Buffer,
// ) (error){

// 	// Create File for the Video on Disk
// 	file_path:=fmt.Sprintf("data/%s/%s",mp4file.data_node_1,mp4file.file_name)
// 	file,err:= os.Create(file_path)

// 	if err !=nil{
// 		return fmt.Errorf("cannot write image to file: %w", err)
// 	}

// 	// Write Data
// 	_, err = mp4Data.WriteTo(file)
// 	if err != nil {
// 		return fmt.Errorf("cannot write image to file: %w", err)
// 	}

// 	// Add File to the lookup table
// 	store.mutex.Lock()
// 	defer store.mutex.Unlock()

// 	store.data[mp4file.file_name]=mp4file
// 	return nil
// }