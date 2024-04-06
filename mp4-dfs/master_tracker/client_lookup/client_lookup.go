package client_

import (
	"fmt"
	"sync"
)

type Client struct {
	socket string
}

func NewClient(socket string) Client {
	return Client{
		socket: socket,
	}
}

type ClientLookUpTable struct {
	mutex sync.RWMutex
	data  map[string]*Client //Map of key are strings and values are Files
}


func NewClientLookUpTable() ClientLookUpTable {
	return ClientLookUpTable{
		data: make(map[string]*Client),
	}
}


//Add New Client
func (store *ClientLookUpTable)AddClient(fileName string,socket string) (error){
	// Add File to the lookup table
	store.mutex.Lock()
	defer store.mutex.Unlock()

	store.data[fileName]=&Client{socket: socket}
	return nil
}
//Remove Client
func (store *ClientLookUpTable)RemoveClient(fileName string) (error){
	store.mutex.Lock()
    defer store.mutex.Unlock()
    delete(store.data, fileName)
	return nil
}

//Add New ClientG
func (store *ClientLookUpTable)GetClientSocket(fileName string) (string){
	_, exists := store.data[fileName]
	if exists {
		return store.data[fileName].socket
	}else {
		return ""
	}
}

//Get Info
func (store *ClientLookUpTable) PrintClientInfo(fileName string )(string){
	client:=store.data[fileName]

	details := fmt.Sprintf("[Client] at %s, Waiting for %s",
	client.socket, fileName)
	return details
}