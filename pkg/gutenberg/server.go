package gutenberg

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/byoc-io/gutenberg/pkg/storage"
	"github.com/byoc-io/gutenberg/pkg/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

const (
	prefix = "/api/v1"
)

// Config holds the server's configuration options.
type Config struct {
	Repository storage.Repository
	Logger     logrus.FieldLogger
}

// Server is the top level object.
type Server struct {
	router     *gin.Engine
	repository storage.Repository
	logger     logrus.FieldLogger
}

// NewServer constructs a server from the provided config.
func NewServer(c Config) (*Server, error) {
	if c.Repository == nil {
		return nil, errors.New("server: storage cannot be null")
	}

	if c.Logger == nil {
		return nil, errors.New("server: logger cannot be null")
	}

	// Router Engine
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(RequestIDMiddleware())
	r.Use(LogMiddleware(c.Logger))

	// Server
	s := &Server{
		router:     r,
		repository: c.Repository,
		logger:     c.Logger,
	}

	// API
	v1 := r.Group(prefix)
	// Projects
	v1.GET("/posts", s.ListPostsHandler)
	v1.GET("/posts/:id", s.GetPostHandler)
	v1.POST("/posts", s.CreatePostHandler)
	v1.PUT("/posts/:id", s.UpdatePostHandler)

	// Base endpoints
	r.GET("/health", s.HealthHandler)
	r.GET("/status", s.StatusHandler)
	r.GET("/", rootHandler)

	return s, nil
}

// RunHTTPServer will serve the http.Server and will monitor for signals
// allowing for graceful termination (SIGTERM) or restart (SIGUSR2).
func (s *Server) RunHTTPServer(addr string) error {
	//return gracehttp.Serve(&http.Server{
	//	Addr:    addr,
	//	Handler: s.router,
	//})

	return s.router.Run(addr)
}

type HealthCheck struct {
	ID     string `json:"id"`
	Node   string `json:"node"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Notes  string `json:"notes"`
	Output string `json:"output"`
}

func NewHealthCheck(name, status, notes, output string) *HealthCheck {
	return &HealthCheck{
		ID:     util.NewUUID(),
		Node:   GetLocalIP(),
		Name:   name,
		Status: status,
		Notes:  notes,
		Output: output,
	}
}

// HealthHandler is gin handle for check system health.
func (s *Server) HealthHandler(c *gin.Context) {
	err := s.repository.HealthCheck()
	if err != nil {
		abortWithError(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	checks := []*HealthCheck{
		NewHealthCheck("storage", "passing", "", ""),
	}

	c.JSON(http.StatusOK, checks)
}

// StatusHandler is gin handle for get system status.
func (s *Server) StatusHandler(c *gin.Context) {
	c.JSON(http.StatusOK, util.GetStats())
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Gutenberg Server",
	})
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
