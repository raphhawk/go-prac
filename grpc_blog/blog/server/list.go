package main

import (
	pb "github.com//raphhawk/grpc-blog/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func ListBlogs(e *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error { return nil }
