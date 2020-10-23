package client

import (
	"sync"
)

type ClientPool struct {
	clients *sync.Map
}

func NewClientPool() *ClientPool {
	return &ClientPool{
		clients: new(sync.Map),
	}
}

func (self *ClientPool) Exist(key string) bool {
	_, hit := self.clients.Load(key)
	return hit
}

func (self *ClientPool) Push(key string, c *Client) {
	self.clients.Store(key, c)
}

func (self *ClientPool) Pull(key string) (c *Client, hit bool) {
	v, hit := self.clients.Load(key)
	if hit {
		self.clients.Delete(key)
		c = v.(*Client)
	}
	return
}

func (self *ClientPool) SendOne(key string, packet []byte) bool {
	v, hit := self.clients.Load(key)
	if hit {
		v.(*Client).Sendch <- packet
	}
	return hit
}

func (self *ClientPool) SendMany(keys []string, packet []byte) {
	for _, k := range keys {
		self.SendOne(k, packet)
	}
}

func (self *ClientPool) SendAll(packet []byte) {
	self.clients.Range(func(k, v interface{}) bool {
		v.(*Client).Sendch <- packet
		return true
	})
}

func (self *ClientPool) CloseOne(key string) bool {
	v, hit := self.clients.Load(key)
	if hit {
		close(v.(*Client).Sendch)
		self.clients.Delete(key)
	}
	return hit
}

func (self *ClientPool) CloseMany(keys []string) {
	for _, k := range keys {
		self.CloseOne(k)
	}
}

func (self *ClientPool) CloseAll() {
	self.clients.Range(func(k, v interface{}) bool {
		close(v.(*Client).Sendch)
		return true
	})
	self.clients = new(sync.Map)
}

var initacceptedclientpoolonce sync.Once
var acceptedclientpool *ClientPool

func AcceptedClientPoolInstance() *ClientPool {
	initacceptedclientpoolonce.Do(func() {
		acceptedclientpool = NewClientPool()
	})
	return acceptedclientpool
}

var initunacceptedclientpoolonce sync.Once
var unacceptedclientpool *ClientPool

func UnacceptedClientPoolInstance() *ClientPool {
	initunacceptedclientpoolonce.Do(func() {
		unacceptedclientpool = NewClientPool()
	})
	return unacceptedclientpool
}

func ClientPoolUnacceptedToAccepted(old, new string) *Client {
	c, hit := UnacceptedClientPoolInstance().Pull(old)
	if hit {
		c.Accepted(true)
		AcceptedClientPoolInstance().Push(new, c)
	}
	return c
}
