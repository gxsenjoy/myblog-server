package myblog

import (
	"sync"

	mgo "gopkg.in/mgo.v2"
)

var (
	db   Database
	once sync.Once
)

type Database interface {
	Connect(url string) error
	Disconnect() error

	SetDB(dbName string) (DB, error)
}

type DB interface {
	SetTable(tblName string) (Table, error)
}

type Table interface {
	Find(query interface{}) Result
	FindOne(id interface{}) Result
}

type Result interface {
	One(model Model) error
}

func NewDatabase() Database {
	once.Do(func() {
		db = newMongodb()
	})
	return db
}

// -- MongoDB --

type mongodb struct {
	session *mgo.Session
}

func (m *mongodb) Connect(url string) (err error) {
	m.session, err = mgo.Dial(url)
	return
}

func (m *mongodb) Disconnect() error {
	m.session.Close()
	return nil
}

func (m *mongodb) SetDB(dbName string) (DB, error) {
	return &mgDB{m.session.DB(dbName)}, nil
}

func newMongodb() *mongodb {
	return &mongodb{}
}

// -- mgDB --

type mgDB struct {
	*mgo.Database
}

func (m *mgDB) SetTable(tblName string) (Table, error) {
	return &mgTable{m.C(tblName)}, nil
}

// -- mgTable --

type mgTable struct {
	*mgo.Collection
}

func (m *mgTable) Find(query interface{}) Result {
	return &mgResult{m.Collection.Find(query)}
}

func (m *mgTable) FindOne(id interface{}) Result {
	return &mgResult{m.Collection.FindId(id)}
}

// -- mgResult --

type mgResult struct {
	*mgo.Query
}

func (m *mgResult) One(model Model) error {
	return m.Query.One(model)
}
