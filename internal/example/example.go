package example

import (
	"strconv"

	"github.com/wangrenjun/lessismore/internal/example/controllers"

	"github.com/wangrenjun/lessismore/pkg/codes"

	"github.com/wangrenjun/lessismore/internal/client"
	"github.com/wangrenjun/lessismore/internal/pack"
)

func init() {
	client.PathRouterInstance().HandleFunc(MyPath, dispatch)

	exampleRouter = client.NewIdRouter()
	exampleRouter.HandleFunc(10, controllers.Login)
	exampleRouter.HandleFunc(11, controllers.Echo)
	exampleRouter.HandleFunc(12, controllers.Getsession)
	exampleRouter.HandleFunc(13, controllers.Send)
	exampleRouter.HandleFunc(14, controllers.Multicast)
	exampleRouter.HandleFunc(15, controllers.Broadcast)
}

var MyPath = "/example"

func dispatch(c *client.Client, packet []byte) bool {
	mid, ok := pack.Field(packet, "MessageId")
	if !ok {
		rep, _ := pack.PackReply(MyPath, codes.RC_MALFORMED_MESSAGE, nil)
		c.Sendch <- rep
		return true
	}
	msgid, err := strconv.Atoi(mid)
	if err != nil {
		rep, _ := pack.PackReply(MyPath, codes.RC_MESSAGE_UNDEFINED, nil)
		c.Sendch <- rep
		return true
	}
	if !exampleRouter.Dispatch(msgid, c, packet) {
		rep, _ := pack.PackReply(MyPath, codes.RC_MESSAGE_UNDEFINED, nil)
		c.Sendch <- rep
		return true
	}
	return true
}

var exampleRouter *client.IdRouter
