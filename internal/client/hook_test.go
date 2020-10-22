package client_test

import (
	"testing"

	"github.com/wangrenjun/lessismore/internal/client"
)

var ttt *testing.T

func aa(c *client.Client) {
	ttt.Log("aa")
}

func bb(c *client.Client) {
	ttt.Log("bb")
}

func cc(c *client.Client) {
	ttt.Log("cc")
}

func dd(c *client.Client) {
	ttt.Log("dd")
}

func TestHook(t *testing.T) {
	ttt = t
	h := client.NewHook()
	h.Hooking(aa)
	h.Hooking(bb)
	h.Hooking(cc)
	h.Hooking(dd)
	h.Trigger(&client.Client{})
	t.Log("done")
}
