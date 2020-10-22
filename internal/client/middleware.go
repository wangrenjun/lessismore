package client

type MiddleWareChain []MiddleWare
type MiddleWare func(WsHandlerFunc) WsHandlerFunc

func (self *PathRouter) MiddleWareUse(handler WsHandlerFunc) {
	self.middlewares = append(self.middlewares, func(WsHandlerFunc) WsHandlerFunc {
		return handler
	})
}

func (self *IdRouter) MiddleWareUse(handler WsHandlerFunc) {
	self.middlewares = append(self.middlewares, func(WsHandlerFunc) WsHandlerFunc {
		return handler
	})
}

func middlewareWrapped(h WsHandlerFunc, m ...MiddleWare) WsHandlerFunc {
	if len(m) < 1 {
		return h
	}
	wrapped := h
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}
	return wrapped
}
