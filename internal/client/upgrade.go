package client

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/wangrenjun/lessismore/internal/config"
	"github.com/wangrenjun/lessismore/internal/log"
)

func init() {
	HttpRouterInstance().GET("/ws", Upgrade)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  config.WS_READ_BUF_SIZE,
	WriteBufferSize: config.WS_WRITE_BUF_SIZE,
}

func beforeClose(c *Client) {
	return
}

func Upgrade(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.LoggerInstance().Error().Err(err).Msg("Upgrade failed")
		return
	}
	cli := NewClient(conn)
	cli.Accepted(false)
	cli.SetSession(NewTemporarySession(conn.RemoteAddr().String()))
	cli.BeforeClose.Hooking(beforeClose)
	UnacceptedClientPoolInstance().Push(cli.Session().Id(), cli)
	cli.Logger().Info().Msg("Client upgraded")
	go cli.Recv()
	go cli.Send()
}
