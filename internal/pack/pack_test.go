package pack_test

import (
	"testing"

	"github.com/wangrenjun/lessismore/internal/pack"

	"github.com/wangrenjun/lessismore/pkg/codes"
)

func TestPack(t *testing.T) {
	h := `{
"Path": "/example",
"MessageId": 880,
"From": "yule"
}`
	path, found := pack.Field([]byte(h), "Path")
	if !found {
		t.Fatal("Not found")
	}
	t.Logf("path: %s", path)

	rep, err := pack.PackReply("/example2", codes.RC_MALFORMED_MESSAGE, nil)
	if err != nil {
		t.Fatalf("PackReply: %v", err)
	}
	t.Logf("rep: %s", string(rep))

	v := struct {
		UserId int
		Name   string
		RoomId int
	}{
		123,
		"bullshit",
		9988,
	}
	rep, err = pack.PackReply("/example3", codes.RC_USER_NOT_FOUND, v)
	if err != nil {
		t.Fatalf("PackReply: %v", err)
	}
	t.Logf("rep: %s", string(rep))
}
