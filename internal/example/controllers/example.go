package controllers

import (
	"encoding/json"
	"time"

	"github.com/wangrenjun/lessismore/internal/config"

	"github.com/wangrenjun/lessismore/internal/client"
	"github.com/wangrenjun/lessismore/internal/pack"
	"github.com/wangrenjun/lessismore/pkg/codes"
)

var MyPath = "/example"

func beforeClose(c *client.Client) {
	c.Logger().Info().Msg("This client will close")
}

func Login(c *client.Client, packet []byte) bool {
	if c.IsAccepted() {
		rep, _ := pack.PackReply(MyPath, codes.RC_OK, nil)
		c.Sendch <- rep
		return true
	}
	req := &struct {
		Path      string
		MessageId string
		UserId    string
		UserName  string
		Photo     string
		Token     string
	}{}
	if err := json.Unmarshal(packet, &req); err != nil {
		rep, _ := pack.PackReply(MyPath, codes.RC_MALFORMED_MESSAGE, nil)
		c.Sendch <- rep
		return true
	}
	if req.Token != "HELLO" {
		rep, _ := pack.PackReply(MyPath, codes.RC_TOKEN_MISMATCH, nil)
		c.Sendch <- rep
		return true
	}
	old := c.Session().Id()
	new := req.UserId
	if client.ClientPoolUnacceptedToAccepted(old, new) != c {
		if config.Configs.DeployEnv == "dev" {
			c.Logger().Panic().Msg("WTF!")
		}
		rep, _ := pack.PackReply(MyPath, codes.RC_SERVER_ERROR, nil)
		c.Sendch <- rep
		return true
	}
	ss := NewExampleSession(new)
	ss.UserId = req.UserId
	ss.UserName = req.UserName
	ss.Photo = req.Photo
	c.SetSession(ss)
	c.BeforeClose.Hooking(beforeClose)
	rsp := struct {
		Path      string
		MessageId string
		UserId    string
		UserName  string
		Photo     string
	}{
		req.Path,
		req.MessageId,
		req.UserId,
		req.UserName,
		req.Photo,
	}
	resp, _ := json.Marshal(rsp)
	c.Sendch <- resp
	return true
}

func Echo(c *client.Client, packet []byte) bool {
	req := &struct {
		Path      string
		MessageId string
		UserId    string
		Message   string
	}{}
	if err := json.Unmarshal(packet, &req); err != nil {
		rep, _ := pack.PackReply(MyPath, codes.RC_MALFORMED_MESSAGE, nil)
		c.Sendch <- rep
		return true
	}
	rsp := &struct {
		Path      string
		MessageId string
		UserId    string
		Message   string
	}{
		req.Path,
		req.MessageId,
		req.UserId,
		req.MessageId,
	}
	resp, _ := json.Marshal(rsp)
	c.Sendch <- resp
	return true
}

func Getsession(c *client.Client, packet []byte) bool {
	resp, _ := json.Marshal(c.Session().(ExampleSession))
	c.Sendch <- resp
	return true
}

func Send(c *client.Client, packet []byte) bool {
	req := &struct {
		Path       string
		MessageId  string
		FromUserId string
		ToUserId   string
		Message    string
	}{}
	if err := json.Unmarshal(packet, &req); err != nil {
		rep, _ := pack.PackReply(MyPath, codes.RC_MALFORMED_MESSAGE, nil)
		c.Sendch <- rep
		return true
	}
	client.AcceptedClientPoolInstance().SendOne(req.ToUserId, packet)
	return true
}

func Multicast(c *client.Client, packet []byte) bool {
	req := &struct {
		Path       string
		MessageId  string
		FromUserId string
		ToUserIds  []string
		Message    string
	}{}
	if err := json.Unmarshal(packet, &req); err != nil {
		rep, _ := pack.PackReply(MyPath, codes.RC_MALFORMED_MESSAGE, nil)
		c.Sendch <- rep
		return true
	}
	client.AcceptedClientPoolInstance().SendMany(req.ToUserIds, packet)
	return true
}

func Broadcast(c *client.Client, packet []byte) bool {
	req := &struct {
		Path       string
		MessageId  string
		FromUserId string
		Message    string
	}{}
	if err := json.Unmarshal(packet, &req); err != nil {
		rep, _ := pack.PackReply(MyPath, codes.RC_MALFORMED_MESSAGE, nil)
		c.Sendch <- rep
		return true
	}
	client.AcceptedClientPoolInstance().SendAll(packet)
	return true
}

type ExampleSession struct {
	id       string
	created  time.Time
	UserId   string
	UserName string
	Photo    string
}

func NewExampleSession(id string) *ExampleSession {
	return &ExampleSession{id: id, created: time.Now()}
}

func (self ExampleSession) Id() string {
	return self.id
}

func (self ExampleSession) Created() time.Time {
	return self.created
}
