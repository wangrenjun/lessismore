package client

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/wangrenjun/lessismore/internal/pack"

	"github.com/rs/zerolog"

	"github.com/gorilla/websocket"
	"github.com/wangrenjun/lessismore/internal/config"
	"github.com/wangrenjun/lessismore/internal/log"
	"github.com/wangrenjun/lessismore/pkg/codes"
)

type Client struct {
	Sendch      chan []byte
	BeforeClose Hook
	conn        *websocket.Conn
	accepted    atomic.Value
	session     atomic.Value
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{Sendch: make(chan []byte, config.CHAN_SEND_BUF_SIZE), BeforeClose: NewHook(), conn: conn}
}

func (c *Client) Accepted(b bool) {
	c.accepted.Store(b)
}

func (c Client) IsAccepted() bool {
	return c.accepted.Load().(bool)
}

func (c *Client) SetSession(session Session) {
	c.session.Store(session)
}

func (c Client) Session() Session {
	ss, ok := c.session.Load().(Session)
	if !ok {
		log.LoggerInstance().Panic().Msg("Session must be implemented")
	}
	return ss
}

func (c Client) String() string {
	return fmt.Sprintf(`{"SessionId": "%s", "Addr": "%s", "Accepted": %v}`, c.Session().Id(), c.conn.RemoteAddr(), c.IsAccepted())
}

func (c Client) Logger() *zerolog.Logger {
	logger := log.LoggerInstance().With().RawJSON("Client", []byte(c.String())).Logger()
	return &logger
}

func (c *Client) Recv() {
	defer func() {
		c.Logger().Info().Msg("Disconnect the client")
		UnacceptedClientPoolInstance().CloseOne(c.Session().Id())
		AcceptedClientPoolInstance().CloseOne(c.Session().Id())
	}()
	c.conn.SetReadLimit(config.WS_READ_LIMIT)
	c.conn.SetReadDeadline(time.Now().Add(config.PONG_WAIT))
	c.conn.SetPongHandler(func(string) error {
		c.Logger().Debug().Msg("Received pong")
		c.conn.SetReadDeadline(time.Now().Add(config.PONG_WAIT))
		return nil
	})
	for {
		_, packet, err := c.conn.ReadMessage()
		if err != nil || len(packet) == 0 {
			c.Logger().Error().Err(err).Msg("ReadMessage error")
			return
		}
		path, found := pack.Field(packet, "Path")
		if !found {
			c.Logger().Info().Msg("Missing Path")
			rep, _ := pack.PackReply(path, codes.RC_MALFORMED_MESSAGE, nil)
			c.Sendch <- rep
			return
		}
		c.Logger().Debug().Int("ReadBytes", len(packet)).Str("Path", path).Msg("Request received")
		if !PathRouterInstance().Dispatch(path, c, packet) {
			c.Logger().Info().Str("Path", path).Msg("Path not found")
			rep, _ := pack.PackReply(path, codes.RC_RESOURCE_NOT_FOUND, nil)
			c.Sendch <- rep
			return
		}
	}
}

func (c *Client) Send() {
	ticker := time.NewTicker(config.PING_PERIOD)
	defer func() {
		ticker.Stop()
		c.Logger().Info().Msg("Disconnect the client")
		UnacceptedClientPoolInstance().CloseOne(c.Session().Id())
		AcceptedClientPoolInstance().CloseOne(c.Session().Id())
		c.BeforeClose.Trigger(c)
		c.conn.Close()
	}()
	for {
		select {
		case packet, ok := <-c.Sendch:
			if !ok {
				return
			}
			c.conn.SetWriteDeadline(time.Now().Add(config.WS_SEND_WAIT))
			if err := c.conn.WriteMessage(websocket.TextMessage, packet); err != nil {
				c.Logger().Error().Err(err).Msg("WriteMessage error")
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(config.WS_SEND_WAIT))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				c.Logger().Error().Err(err).Msg("PingMessage error")
				return
			}
		}
	}
}
