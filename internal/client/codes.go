package client

import (
	_ "expvar"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangrenjun/lessismore/pkg/codes"
)

func init() {
	HttpRouterInstance().GET("/codes", Codes)
}

func Codes(c *gin.Context) {
	rcs := []codes.ReturnCode{}
	for i := 0; i < int(codes.RC_MAXIMUM); i++ {
		rcs = append(rcs, codes.ReturnCode(i))
	}
	c.JSON(http.StatusOK, rcs)
}
