package client

import (
	"net/http"
	"sync"

	"github.com/wangrenjun/lessismore/pkg/codes"

	"github.com/gin-gonic/gin"
)

var inithttprouteronce sync.Once
var router *gin.Engine

func HttpRouterInstance() *gin.Engine {
	inithttprouteronce.Do(func() {
		router = gin.New()
		router.NoRoute(func(c *gin.Context) {
			c.JSON(http.StatusNotFound, gin.H{"Code": codes.RC_RESOURCE_NOT_FOUND})
		})
	})
	return router
}
