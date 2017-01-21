package authoring

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	"github.com/byoc-io/gutenberg/pkg/model"
	kitlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

// MakeHandler returns a handler for the authoring service.
func MakeHandler(ctx context.Context, s Service, logger kitlog.Logger) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(s)
	opts := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	prefix := "/authoring/v1"

	// GET 		/posts/:key		retrieves a post by key (id or slug)
	r.Methods("GET").Path(prefix + "/posts/{key}").Handler(httptransport.NewServer(
		ctx,
		e.GetPostEndpoint,
		decodeGetPostRequest,
		encodeResponse,
		opts...,
	))

	return r
}

func decodeGetPostRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	key, ok := vars["key"]
	if !ok {
		return nil, errors.New("bad route")
	}
	return getPostRequest{Key: key}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	switch err {
	case model.ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	case ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
