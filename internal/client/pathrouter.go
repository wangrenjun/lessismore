package client

import (
	"path/filepath"
	"sync"

	"github.com/wangrenjun/lessismore/internal/log"
)

type WsHandlerFunc func(*Client, []byte) bool

type PathRouter struct {
	handlers    map[string]WsHandlerFunc
	middlewares MiddleWareChain
}

func NewPathRouter() *PathRouter {
	return &PathRouter{handlers: make(map[string]WsHandlerFunc)}
}

func (self *PathRouter) HandleFunc(pattern string, handler WsHandlerFunc) {
	if self.handlers[pattern] == nil {
		self.handlers[pattern] = handler
	} else {
		log.LoggerInstance().Panic().Str("Pattern", pattern).Msg("Pattern already used")
	}
}

func (self PathRouter) find(path string) WsHandlerFunc {
	handler, hit := self.handlers[path]
	if !hit {
		for pattern, h := range self.handlers {
			if matched, err := filepath.Match(pattern, path); matched && err == nil {
				handler = h
				break
			}
		}
	}
	return handler
}

func (self PathRouter) Dispatch(path string, c *Client, packet []byte) bool {
	if handler := self.find(path); handler != nil {
		for i := 0; i < len(self.middlewares)+1; i++ {
			if !middlewareWrapped(handler, self.middlewares[i:]...)(c, packet) {
				return true
			}
		}
		return true
	}
	return false
}

func (self PathRouter) Range(cb func(string, WsHandlerFunc)) {
	for pattern, handler := range self.handlers {
		cb(pattern, handler)
	}
}

var initpathrouteronce sync.Once
var pathrouter *PathRouter

func PathRouterInstance() *PathRouter {
	initpathrouteronce.Do(func() {
		pathrouter = NewPathRouter()
	})
	return pathrouter
}
