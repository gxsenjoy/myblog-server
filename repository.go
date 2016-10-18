package myblog

import "errors"

type Repository interface {
	Table
}

func NewRepository(dbName, colName string) (Repository, error) {
	switch colName {
	case "posts":
		return newPostRepository(dbName)
	default:
		return nil, errors.New("An error has occurred")
	}
}

// -- postRepository --

type postRepository struct {
	Table
}

func newPostRepository(dbName string) (*postRepository, error) {
	db, err := db.SetDB(dbName)
	if err != nil {
		return nil, err
	}

	tbl, err := db.SetTable("posts")
	if err != nil {
		return nil, err
	}

	return &postRepository{tbl}, nil
}
