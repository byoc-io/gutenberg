package database

import (
	"github.com/byoc-io/gutenberg/pkg/storage"
	"gopkg.in/mgo.v2"
)

const (
	PostsCollection = "posts"
	UsersCollection = "users"
)

// repository is a Mongodb implementation of the Repository.
type repository struct {
	session *mgo.Session
	name    string
}

func New(db *mgo.Database) storage.Repository {
	return &repository{
		session: db.Session,
		name:    db.Name,
	}
}

func EnsureIndex(db *mgo.Database) error {
	session := db.Session.Copy()
	defer session.Close()

	posts := session.DB(db.Name).C(PostsCollection)

	id := mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	if err := posts.EnsureIndex(id); err != nil {
		return err
	}

	slug := mgo.Index{
		Key:        []string{"slug"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	if err := posts.EnsureIndex(slug); err != nil {
		return err
	}

	return nil
}

func (r *repository) HealthCheck() error {
	session := r.session.Copy()
	defer session.Close()

	return session.Ping()
}
