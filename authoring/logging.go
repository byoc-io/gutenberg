package authoring

import (
	"github.com/byoc-io/gutenberg/pkg/model"
	"github.com/go-kit/kit/log"
	"time"
)

type loggingService struct {
	logger log.Logger
	Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) GetPost(key string) (p model.Post, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GET",
			"endpoint", "/posts",
			"key", key,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetPost(key)
}
