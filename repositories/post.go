package repositories

import (
	myblog "github.com/nomkhonwaan/myblog-server"
	"github.com/nomkhonwaan/myblog-server/models"
)

func NewPostRepository(db myblog.Database) Repository {
	return &postRepository{db.Collection("posts")}
}

type postRepository struct {
	myblog.Collection
}

func (p *postRepository) Create(m models.Model) (models.Model, error) {
	return nil, nil
}

func (p *postRepository) Find(query interface{}) ([]models.Model, error) {
	var entities []models.Model
	err := p.All(query.(myblog.Query))(entities)

	return entities, err
}

func (p *postRepository) FindByID(id interface{}) (models.Model, error) {
	entity := models.NewPostModel()
	err := p.First(myblog.Query{"id": id})(entity)

	return entity, err
}

func (p *postRepository) Update(id interface{}, m models.Model) error {
	return nil
}

func (p *postRepository) Delete(id interface{}) error {
	return nil
}
