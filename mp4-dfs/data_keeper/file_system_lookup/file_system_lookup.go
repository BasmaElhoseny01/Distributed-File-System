package file_system_lookup

import (
	"fmt"
	"sync"
)

type File struct {
	// Define Fields For File
	file_name  string
	path string
}

func NewFile(file_name string, path string) File {
	return File{
		file_name:   file_name,
		path: path,
	}
}

type FileSystemLookUpTable struct {
	mutex sync.RWMutex
	data  map[string]*File //Map of key are strings and values are Files
}


func NewFileSystemLookUpTable() FileSystemLookUpTable {
	return FileSystemLookUpTable{
		data: make(map[string]*File),
	}
}

//Add New File
func (store *FileSystemLookUpTable)AddFile(mp4file *File) (error){
	// Add File to the lookup table
	store.mutex.Lock()
	defer store.mutex.Unlock()

	store.data[mp4file.file_name]=mp4file
	for key, value := range store.data {
		fmt.Printf("[File System] Key: %s, Value: %v\n", key, value)
	}
	return nil
}


//Get Info
func (store *FileSystemLookUpTable) PrintFileInfo(fileName string)(string){
	file:=store.data[fileName]

	details := fmt.Sprintf("[File System] Name: %s, at %s",file.file_name,file.path)
	return details
}

// Get File Path
func (store *FileSystemLookUpTable) GetFilePath(fileName string)(string){
	file:=store.data[fileName]

	return file.path
}