package database

import (
	"fmt"
	"github.com/byoc-io/gutenberg/pkg/model"
	"github.com/byoc-io/gutenberg/pkg/storage"
	"github.com/byoc-io/gutenberg/pkg/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func (r *repository) ListPosts(pagination *storage.Pagination) ([]*model.Post, error) {
	session := r.session.Copy()
	defer session.Close()

	c := session.DB(r.name).C(PostsCollection)

	var posts []*model.Post
	err := c.Find(bson.M{}).All(&posts)
	if err != nil {
		return nil, err
	}

	if posts == nil {
		posts = []*model.Post{}
	}

	return posts, nil
}

func (r *repository) GetPost(key string) (*model.Post, error) {
	session := r.session.Copy()
	defer session.Close()

	c := session.DB(r.name).C(PostsCollection)

	var post model.Post
	if err := c.Find(bson.M{"$or": []bson.M{bson.M{"id": key}, bson.M{"slug": key}}}).One(&post); err != nil {
		if err == mgo.ErrNotFound {
			return nil, storage.ErrNotFound
		}

		return nil, err
	}

	return &post, nil
}

func (r *repository) InsertPost(post *model.Post) error {
	now := time.Now()
	if post.ID == "" {
		post.ID = util.NewUUID()
	}

	post.CreatedAt = &now
	post.UpdatedAt = &now

	session := r.session.Copy()
	defer session.Close()

	c := session.DB(r.name).C(PostsCollection)

	if err := c.Insert(post); err != nil {
		if mgo.IsDup(err) {
			return storage.ErrAlreadyExists
		}

		return err
	}

	return nil
}

func (r *repository) UpdatePost(post *model.Post) error {
	now := time.Now()
	if post.ID == "" {
		return fmt.Errorf("invalid post, id cannot be empty")
	}

	post.UpdatedAt = &now

	session := r.session.Copy()
	defer session.Close()

	c := session.DB(r.name).C(PostsCollection)

	if err := c.Update(bson.M{"id": post.ID}, post); err != nil {
		if err == mgo.ErrNotFound {
			return storage.ErrNotFound
		}

		return err
	}

	return nil
}
