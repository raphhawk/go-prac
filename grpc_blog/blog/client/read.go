package main

import (
	"context"
	"log"

	pb "github.com/raphhawk/grpc-blog/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("---Invoked readBlog function---")
	req := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Printf("Error while reading: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)
	return res
}
