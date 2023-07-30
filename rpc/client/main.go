package main

import (
	"log"
	"net/rpc"
)

// Item: Exportable struct
type Item struct {
	Title string
	Body  string
}

func main() {
	reply := Item{}

	// communicate to RPC server over http
	client, err := rpc.DialHTTP("unix", "/var/mapreduce")
	if err != nil {
		log.Fatalf("Error Connecting to Server %v\n", err)
	}

	// perform CRUD operations over DB
	client.Call("DB.CreateItem", Item{"1", "1body"}, &reply)
	log.Printf("Reply of read: %v\n", reply)
	client.Call("DB.CreateItem", Item{"2", "2body"}, &reply)
	client.Call("DB.CreateItem", Item{"3", "3body"}, &reply)
	client.Call("DB.UpdateItem", []string{"2", "2nf", "2nfbody"}, &reply)
	client.Call("DB.DeleteItem", Item{"3", "3body"}, &reply)
}
