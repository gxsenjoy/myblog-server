package repositories

import "github.com/nomkhonwaan/myblog-server/models"

type Repository interface {
	Create(m models.Model) (models.Model, error)
	Find(query interface{}) ([]models.Model, error)
	FindByID(id interface{}) (models.Model, error)
	Update(id interface{}, m models.Model) error
	Delete(id interface{}) error
}
