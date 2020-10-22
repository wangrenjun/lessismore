package example

import "time"

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
