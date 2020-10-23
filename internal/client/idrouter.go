package client

import (
	"sync"

	"github.com/wangrenjun/lessismore/internal/log"
)

type IdRouter struct {
	handlers    map[int]WsHandlerFunc
	middlewares MiddleWareChain
}

func NewIdRouter() *IdRouter {
	return &IdRouter{handlers: make(map[int]WsHandlerFunc)}
}

func (self *IdRouter) HandleFunc(id int, handler WsHandlerFunc) {
	if self.handlers[id] == nil {
		self.handlers[id] = handler
	} else {
		log.LoggerInstance().Panic().Int("Id", id).Msg("Id already used")
	}
}

func (self IdRouter) Dispatch(id int, c *Client, packet []byte) bool {
	if handler := self.handlers[id]; handler != nil {
		for i := 0; i < len(self.middlewares)+1; i++ {
			if !middlewareWrapped(handler, self.middlewares[i:]...)(c, packet) {
				return true
			}
		}
		return true
	}
	return false
}

func (self IdRouter) Range(cb func(int, WsHandlerFunc)) {
	for id, handler := range self.handlers {
		cb(id, handler)
	}
}

var initidrouteronce sync.Once
var idrouter *IdRouter

func IdRouterInstance() *IdRouter {
	initidrouteronce.Do(func() {
		idrouter = NewIdRouter()
	})
	return idrouter
}
