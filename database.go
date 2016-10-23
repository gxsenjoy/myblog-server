package myblog

import (
	"sync"

	"github.com/golang/glog"
	"github.com/nomkhonwaan/myblog-server/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	conn Connection
	once sync.Once
)

func NewConnection() Connection {
	once.Do(func() {
		conn = &mgConnection{}
	})
	return conn
}

type Connection interface {
	Connect(url string) error
	Disconnect() error

	Database(dbName string) Database
}

type Database interface {
	Collection(colName string) Collection
}

type Collection interface {
	All(q Query) func(m []models.Model) error
	First(q Query) func(m models.Model) error
}

type Query map[string]interface{}

type mgConnection struct {
	*mgo.Session
}

func (mg *mgConnection) Connect(url string) (err error) {
	mg.Session, err = mgo.Dial(url)
	return err
}

func (mg *mgConnection) Disconnect() error {
	mg.Session.Close()
	return nil
}

func (mg *mgConnection) Database(dbName string) Database {
	return &mgDB{mg.DB(dbName)}
}

type mgDB struct {
	*mgo.Database
}

func (mg *mgDB) Collection(colName string) Collection {
	return &mgCollection{mg.C(colName)}
}

type mgCollection struct {
	*mgo.Collection
}

func (mg *mgCollection) First(q Query) func(m models.Model) error {
	return func(m models.Model) error {
		func() {
			defer func() {
				if r := recover(); r != nil {
					glog.Warning("An error has occurred: %v", r)
				}
			}()
			if _, ok := q["id"]; ok {
				q["_id"] = bson.ObjectIdHex(q["id"].(string))
				delete(q, "id")
			}
		}()

		return mg.Find(q).One(m)
	}
}

func (mg *mgCollection) All(q Query) func(m []models.Model) error {
	return func(m []models.Model) error {
		if _, ok := q["id"]; ok {
			delete(q, "id")
		}

		return mg.Find(q).All(&m)
	}
}
