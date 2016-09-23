package main

import (
	"flag"
	"net"

	"github.com/golang/glog"
	"github.com/nomkhonwaan/myblog-server/blog-service/posts"
	"github.com/nomkhonwaan/myblog-server/protos/blog-service/posts"
	"google.golang.org/grpc"
)

func Run() error {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	protos_blogservice_posts.RegisterPostsServiceServer(s, &blogservice_posts.PostsServiceServer{})

	s.Serve(lis)
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := Run(); err != nil {
		glog.Fatal(err)
	}
}
