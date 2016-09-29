package blogservice_posts

import (
	"golang.org/x/net/context"

	pb "github.com/nomkhonwaan/myblog-server/protos/blog-service/posts"
)

type PostsServiceServer struct{}

func (s *PostsServiceServer) FindByID(ctx context.Context, in *pb.PostIDRequest) (*pb.Post, error) {
	return &pb.Post{
		PostTitle: "[gRPC] รู้จักกับ gRPC โพโตคอลเก่าเอามาปัดฝุ่นใหม่โดย Google",
		PostSlug:  "grpc-getting-started-with-grpc-by-google",
	}, nil
}
