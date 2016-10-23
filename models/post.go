package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

func NewPostModel() Model {
	return &mgPost{}
}

type Post interface {
	Model
	GetTitle() string
	GetSlug() string
	GetPublishedAt() time.Time
	GetMarkdown() string
	GetHTML() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type post struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	PublishedAt time.Time `json:"publishedAt"`
	Markdown    string    `json:"markdown"`
	HTML        string    `json:"html"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (p *post) GetID() string {
	return p.ID
}

func (p *post) GetTitle() string {
	return p.Title
}

func (p *post) GetSlug() string {
	return p.Slug
}

func (p *post) GetPublishedAt() time.Time {
	return p.PublishedAt
}

func (p *post) GetMarkdown() string {
	return p.Markdown
}

func (p *post) GetHTML() string {
	return p.HTML
}

func (p *post) GetCreatedAt() time.Time {
	return p.CreatedAt
}

func (p *post) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

type mgPost struct {
	post `bson:",inline"`
	ID   bson.ObjectId `json:"id" bson:"_id"`
}

func (p *mgPost) GetID() string {
	return p.ID.Hex()
}
