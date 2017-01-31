package gutenberg

import (
	"github.com/byoc-io/gutenberg/pkg/storage"
	"github.com/gin-gonic/gin"
	"strconv"
)

func abortWithError(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func getPagination(c *gin.Context) *storage.Pagination {
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", strconv.Itoa(storage.Limit)))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	return storage.NewPagination(perPage, page)
}
