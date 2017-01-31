package gutenberg

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/byoc-io/gutenberg/pkg/util"
	"github.com/byoc-io/gutenberg/version"
	"github.com/gin-gonic/gin"
)

// LogMiddleware provide gin router handler.
func LogMiddleware(logger logrus.FieldLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		log := &logReq{
			URI:         c.Request.URL.Path,
			Method:      c.Request.Method,
			IP:          c.ClientIP(),
			ContentType: c.ContentType(),
			Agent:       c.Request.Header.Get("User-Agent"),
		}

		// format is string
		output := fmt.Sprintf("%s %s %s %s %s",
			log.Method,
			log.URI,
			log.IP,
			log.ContentType,
			log.Agent,
		)

		// TODO: Use logger
		logger.Debug(output)

		c.Next()
	}
}

// RequestIDMiddleware injects a special header X-Request-Id to response headers
// that could be used to track incoming requests for monitoring/debugging
// purposes.
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", util.NewUUID())
		c.Next()
	}
}

// VersionMiddleware : add version on header.
func VersionMiddleware() gin.HandlerFunc {
	// Set out header value for each response
	return func(c *gin.Context) {
		c.Header("X-Revision", version.Version)
		c.Next()
	}
}
