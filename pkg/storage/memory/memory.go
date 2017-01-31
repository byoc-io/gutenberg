package memory

import (
	"github.com/byoc-io/gutenberg/pkg/model"
	"github.com/byoc-io/gutenberg/pkg/storage"
	"github.com/byoc-io/gutenberg/pkg/util"
	"sync"
)

type repository struct {
	mtx   sync.RWMutex
	posts map[string]*model.Post
}

func (r *repository) ListPosts(pagination *storage.Pagination) ([]*model.Post, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	posts := []*model.Post{}
	for _, p := range r.posts {
		posts = append(posts, p)
	}

	return posts, nil
}

func (r *repository) GetPost(key string) (*model.Post, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if val, ok := r.posts[key]; ok {
		return val, nil
	}

	return nil, storage.ErrNotFound
}

func (r *repository) InsertPost(post *model.Post) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	id := util.NewUUID()
	post.ID = id
	r.posts[id] = post
	return nil
}

func (r *repository) UpdatePost(post *model.Post) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	post, ok := r.posts[post.ID]
	if !ok {
		return storage.ErrNotFound
	}

	r.posts[post.ID] = post
	return nil
}

func (r *repository) HealthCheck() error {
	return nil
}

func (r *repository) String() string {
	return "in-memory storage"
}

// NewRepository returns an in memory repository.
func NewRepository() storage.Repository {
	posts := make(map[string]*model.Post)
	return &repository{
		posts: posts,
	}
}
