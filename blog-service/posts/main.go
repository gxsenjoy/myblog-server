package blogservice_posts

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/golang/glog"
	myblog "github.com/nomkhonwaan/myblog-server"
	pb "github.com/nomkhonwaan/myblog-server/protos/blog-service/posts"
	"golang.org/x/net/context"
)

type PostsServiceServer struct {
	myblog.Repository
}

func (s *PostsServiceServer) FindByID(ctx context.Context, in *pb.PostIDRequest) (*pb.Post, error) {
	var postModel myblog.PostModel
	var err error

	func() {
		defer func() {
			if r := recover(); r != nil {
				glog.Info("An error has occurred: %v", r)
				err = s.Find(bson.M{"slug": in.Id}).One(&postModel)
			}
		}()
		err = s.FindOne(bson.ObjectIdHex(in.Id)).One(&postModel)
	}()

	return &pb.Post{
		Id:        postModel.ID.Hex(),
		PostTitle: postModel.PostTitle,
		PostSlug:  postModel.PostSlug,
		// PublishedAt: postModel.PublishedAt,
		Html: postModel.HTML,
		// CreatedAt: postModel.CreatedAt,
		// UpdatedAt: postModel.UpdatedAt,
	}, err
}
