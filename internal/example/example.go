package example

import (
	"strconv"
	"sync"

	"github.com/wangrenjun/lessismore/pkg/codes"

	"github.com/wangrenjun/lessismore/internal/client"
	"github.com/wangrenjun/lessismore/internal/pack"
)

func init() {
	client.PathRouterInstance().HandleFunc(MyPath, dispatch)
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
	if !IdRouterInstance().Dispatch(msgid, c, packet) {
		rep, _ := pack.PackReply(MyPath, codes.RC_MESSAGE_UNDEFINED, nil)
		c.Sendch <- rep
		return true
	}
	return true
}

var initexampleidrouteronce sync.Once
var exampleidrouter *client.IdRouter

func IdRouterInstance() *client.IdRouter {
	initexampleidrouteronce.Do(func() {
		exampleidrouter = client.NewIdRouter()
	})
	return exampleidrouter
}
