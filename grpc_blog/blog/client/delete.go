package main

import (
	"context"
	"log"

	pb "github.com/raphhawk/grpc-blog/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("---Invoked deleteBlog function---")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("Error while deleting: %v\n", err)
	}

	log.Println("Blog was deleted")

}
