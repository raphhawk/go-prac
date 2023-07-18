package main

import (
	"context"
	"log"

	pb "github.com/raphhawk/grpc-blog/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("---Creating blog---")

	blog := &pb.Blog{
		AuthorId: "Alan Moore",
		Title:    "The Swamp Thing Vol 26",
		Content:  "The Swamp Realised he is not a Human",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Error Creating blog: %v\n", err)
	}

	log.Println("Blog has been created with Id: " + res.Id)
	return res.Id
}
