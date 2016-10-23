package main

import (
	"flag"
	"net"
	"os"

	"github.com/golang/glog"
	myblog "github.com/nomkhonwaan/myblog-server"
	"github.com/nomkhonwaan/myblog-server/blog-service/posts"
	"github.com/nomkhonwaan/myblog-server/protos/blog-service/posts"
	"github.com/nomkhonwaan/myblog-server/repositories"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var (
	conn myblog.Connection
)

func Run(port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	s := grpc.NewServer()

	protos_blogservice_posts.RegisterPostsServiceServer(s, &blogservice_posts.PostsServiceServer{
		repositories.NewPostRepository(
			conn.Database(viper.GetString("databases.mongodb.dbname")),
		),
	})

	s.Serve(lis)
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := myblog.InitConfig(myblog.Config{
		"services.server.port":     os.Getenv("SERVICES_SERVER_PORT"),
		"databases.mongodb.url":    os.Getenv("DATABASES_MONGODB_URL"),
		"databases.mongodb.dbname": os.Getenv("DATABASES_MONGODB_DBNAME"),
	}); err != nil {
		glog.Fatal(err)
	}

	conn = myblog.NewConnection()
	defer conn.Disconnect()

	if err := conn.Connect(viper.GetString("databases.mongodb.url")); err != nil {
		glog.Fatal(err)
	}

	if err := Run(viper.GetString("services.server.port")); err != nil {
		glog.Fatal(err)
	}
}
