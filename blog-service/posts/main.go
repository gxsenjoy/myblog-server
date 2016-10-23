package blogservice_posts

import (
	"errors"

	"github.com/nomkhonwaan/myblog-server/models"
	pb "github.com/nomkhonwaan/myblog-server/protos/blog-service/posts"
	"github.com/nomkhonwaan/myblog-server/repositories"
	"golang.org/x/net/context"
)

type PostsServiceServer struct {
	repositories.Repository
}

func (s *PostsServiceServer) FindByID(ctx context.Context, in *pb.PostIDRequest) (*pb.Post, error) {
	result, err := s.Repository.FindByID(in.Id)

	if entity, ok := result.(models.Post); ok {
		return &pb.Post{
			Id:        entity.GetID(),
			PostTitle: entity.GetTitle(),
			PostSlug:  entity.GetSlug(),
			// PublishedAt: entity.GetPublishedAt(),
			Html: entity.GetHTML(),
			// CreatedAt:   entity.GetCreatedAt(),
			// UpdatedAt:   entity.GetUpdatedAt(),
		}, err
	}

	return &pb.Post{}, errors.New("An error has occurred.")
}
