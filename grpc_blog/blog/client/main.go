package main

import (
	"log"

	pb "github.com/raphhawk/grpc-blog/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)
	id := createBlog(c)
	readBlog(c, id)
	readBlog(c, "23l4k2345lk3456")
	updateBlog(c, id)
	listBlog(c)
	deleteBlog(c, id)
}
