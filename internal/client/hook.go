package client

type Hook chan func(c *Client)

func NewHook() Hook {
	return make(chan func(c *Client), 32)
}

func (self Hook) Hooking(cb func(c *Client)) {
	select {
	case self <- cb:
	default:
	}
}

func (self Hook) Trigger(c *Client) {
	for {
		select {
		case cb := <-self:
			cb(c)
		default:
			return
		}
	}
}
