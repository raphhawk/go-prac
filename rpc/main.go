package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// Item: Exportable Database Document Struct
type Item struct {
	Title string
	Body  string
}

// DB: Exportable Database Struct
type DB struct {
	db []Item
}

// CRUD: Exportable Interface for CRUD implementation over DB
type CRUD interface {
	ReadOneItem(req string, resp *Item) error
	CreateItem(req Item, resp *Item) error
	UpdateItem(req []string, resp *Item) error
	DeleteItem(req Item, resp *Item) error
}

// makeDBCrud: Get CRUD instance for rpc communication
func makeDBCrud() CRUD {
	// ensure database implements all necessary functions
	return &DB{}
}

// PrintAll: Prints entire database array locally
func (d *DB) PrintAll() {
	log.Printf("%v\n", d.db)
}

// ReadOneItem: Finds one Item in db that has matching title
func (d *DB) ReadOneItem(req string, resp *Item) error {
	getItem := Item{}
	for _, i := range d.db {
		if i.Title == req {
			getItem = i
		}
	}
	resp = &getItem
	d.PrintAll()
	return nil
}

// CreateItem: Creats one Item in db with given Item
func (d *DB) CreateItem(req Item, resp *Item) error {
	d.db = append(d.db, req)
	resp = &req
	d.PrintAll()
	return nil
}

// UpdateItem: Updates one Item in db with given title
func (d *DB) UpdateItem(req []string, resp *Item) error {
	i := Item{Title: req[1], Body: req[2]}
	for idx, it := range d.db {
		if it.Title == req[0] {
			d.db[idx] = i
			break
		}
	}
	resp = &i
	d.PrintAll()
	return nil
}

// DeleteItem: Deletes one Item in db with given Item description
func (d *DB) DeleteItem(req Item, resp *Item) error {
	for i, c := range d.db {
		if c.Title == req.Title && c.Body == req.Body {
			d.db = append(d.db[:i], d.db[i+1:]...)
			resp = &c
			break
		}
	}
	d.PrintAll()
	return nil
}

func main() {
	// Register CRUD methods for DB
	err := rpc.Register(makeDBCrud())
	if err != nil {
		log.Fatalf("Error Registering Crud Functions %v\n", err)
	}

	// manage RPC messages over http for client
	rpc.HandleHTTP()
	// listen on unix socket
	listener, err := net.Listen("unix", "/var/mapreduce")
	if err != nil {
		log.Fatalf("Error Listening to socket %v\n", err)
	}

	log.Println("Serving RPC on socket /var/mapreduce")
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatalf("Error Serving on socket %v\n", err)
	}
}
