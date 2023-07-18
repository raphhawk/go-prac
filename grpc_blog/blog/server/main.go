package main

import (
	"context"
	"log"
	"net"

	pb "github.com/raphhawk/grpc_blog/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection

type Server struct {
	pb.BlogServiceServer
}

const addr = "localhost:50051"

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalf("Error connecting to mongodb %v\n", err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatalf("Error connecting to mongodb %v\n", err)
	}

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error connecting to port 50051: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	srv := grpc.NewServer()
	pb.RegisterBlogServiceServer(srv, &Server{})

	if err = srv.Serve(lis); err != nil {
		log.Fatalf("Error serving to port 50051: %v\n", err)
	}
}
