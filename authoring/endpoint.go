package authoring

import (
	"golang.org/x/net/context"

	"github.com/byoc-io/gutenberg/pkg/model"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetPostEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetPostEndpoint: MakeGetPostEndpoint(s),
	}
}

type getPostRequest struct {
	Key string
}

type getPostResponse struct {
	Post  *model.Post `json:"post,omitempty"`
	Error error       `json:"error,omitempty"`
}

func (r getPostResponse) error() error {
	return r.Error
}

func MakeGetPostEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getPostRequest)
		p, err := s.GetPost(req.Key)
		return getPostResponse{Post: &p, Error: err}, nil
	}
}
