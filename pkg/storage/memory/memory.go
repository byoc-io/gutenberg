package memory

import (
	"github.com/byoc-io/gutenberg/pkg/model"
	"github.com/byoc-io/gutenberg/pkg/util"
	"sync"
)

// New returns an in memory storage
type repository struct {
	mtx   sync.RWMutex
	posts map[string]*model.Post
}

func (r *repository) GetPost(key string) (*model.Post, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if val, ok := r.posts[key]; ok {
		return val, nil
	}

	return nil, model.ErrUnknown
}

func (r *repository) CreatePost(post *model.Post) (*model.Post, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	id := util.NewUUID()
	post.ID = id
	r.posts[id] = post
	return post, nil
}

func (r *repository) String() string {
	return "in-memory storage"
}

// NewPostRepository
func NewRepository() model.Repository {
	posts := make(map[string]*model.Post)
	return &repository{
		posts: posts,
	}
}
