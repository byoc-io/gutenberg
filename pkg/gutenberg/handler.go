package gutenberg

import (
	"github.com/byoc-io/gutenberg/pkg/model"
	"github.com/byoc-io/gutenberg/pkg/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func (s *Server) ListPostsHandler(c *gin.Context) {
	pagination := getPagination(c)
	posts, err := s.repository.ListPosts(pagination)
	if err != nil {
		s.logger.Errorf("failed to list posts: %v", err)
		abortWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	//t := strconv.Itoa(total)
	//c.Header("X-Total-Count", t)
	c.JSON(http.StatusOK, posts)
}

func (s *Server) GetPostHandler(c *gin.Context) {
	post, err := s.handleGetPost(c)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, post)
}

func (s *Server) handleGetPost(c *gin.Context) (*model.Post, error) {
	id := c.Param("id")
	post, err := s.repository.GetPost(id)
	if err != nil {
		if err == storage.ErrNotFound {
			s.logger.Errorf("post %s not found: %v", id, err)
			abortWithError(c, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return nil, err
		}

		s.logger.Errorf("failed to get post %s: %v", id, err)
		abortWithError(c, http.StatusInternalServerError, err.Error())
		return nil, err
	}

	return post, nil
}

func (s *Server) CreatePostHandler(c *gin.Context) {
	var p model.Post
	if err := c.BindJSON(&p); err != nil {
		s.logger.Errorf("failed to parse post: %v", err)
		abortWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := s.repository.InsertPost(&p); err != nil {
		s.logger.Errorf("failed to insert post: %v", err)

		if err == storage.ErrAlreadyExists {
			abortWithError(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		abortWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	location := path.Join(prefix, "posts", p.ID)
	c.Header("Location", location)
	c.Writer.WriteHeader(http.StatusCreated)
}

func (s *Server) UpdatePostHandler(c *gin.Context) {
	var p model.Post
	if err := c.BindJSON(&p); err != nil {
		s.logger.Errorf("failed to parse post: %v", err)
		abortWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := s.repository.UpdatePost(&p); err != nil {
		s.logger.Errorf("failed to update post: %v", err)
		abortWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
