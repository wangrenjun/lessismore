package client

import "time"

type Session interface {
	Id() string
}

type TemporarySession struct {
	id      string
	created time.Time
}

func NewTemporarySession(id string) *TemporarySession {
	return &TemporarySession{id: id, created: time.Now()}
}

func (self TemporarySession) Id() string {
	return self.id
}

func (self TemporarySession) Created() time.Time {
	return self.created
}
