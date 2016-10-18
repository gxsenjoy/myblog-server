package myblog

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Model interface{}

type PostModel struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	PostTitle   string        `json:"post_title" bson:"title"`
	PostSlug    string        `json:"post_slug" bson:"slug"`
	PublishedAt time.Time     `json:"published_at" bson:"publishedAt"`
	Markdown    string        `json:"markdown" bson:"markdown"`
	HTML        string        `json:"html" bson:"html"`
	CreatedAt   time.Time     `json:"created_at" bson:"createdAt"`
	UpdatedAt   time.Time     `json:"created_at" bson:"updatedAt"`
}
