package main

import (
	"context"
	"log"

	pb "github.com/raphhawk/grpc-blog/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---Invoked updateBlogs function---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Alan Moore",
		Title:    "Not The Swamp Thing",
		Content:  "The Sawmp was a human",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error while updating: %v\n", err)
	}

	log.Println("Blog was updated!")
}
