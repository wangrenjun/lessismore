package client_test

import (
	"testing"

	"github.com/wangrenjun/lessismore/internal/client"
)

func module1handler(c *client.Client, packet []byte) bool {
	tt.Log("module1handler")
	return true
}

func module2handler(c *client.Client, packet []byte) bool {
	tt.Log("module2handler")
	return true
}

func module3handler(c *client.Client, packet []byte) bool {
	tt.Log("module3handler")
	return true
}

func module4handler(c *client.Client, packet []byte) bool {
	tt.Log("module4handler")
	return true
}

func middlewareA(c *client.Client, packet []byte) bool {
	tt.Log("middlewareA")
	return true
}

func middlewareB(c *client.Client, packet []byte) bool {
	tt.Log("middlewareB")
	return true
}

func middlewareC(c *client.Client, packet []byte) bool {
	tt.Log("middlewareC")
	return true
}

var tt *testing.T

func TestMiddleware(t *testing.T) {
	tt = t
	client.PathRouterInstance().MiddleWareUse(middlewareA)
	client.PathRouterInstance().MiddleWareUse(middlewareB)
	client.PathRouterInstance().MiddleWareUse(middlewareC)
	client.PathRouterInstance().HandleFunc("module1", module1handler)
	client.PathRouterInstance().HandleFunc("module2", module2handler)
	client.PathRouterInstance().HandleFunc("module3", module3handler)
	client.PathRouterInstance().HandleFunc("module*", module4handler)
	if !client.PathRouterInstance().Dispatch("module1", nil, nil) {
		t.Fatal("Dispatch")
	}
	if !client.PathRouterInstance().Dispatch("module2", nil, nil) {
		t.Fatal("Dispatch")
	}
	if !client.PathRouterInstance().Dispatch("module3", nil, nil) {
		t.Fatal("Dispatch")
	}
	if !client.PathRouterInstance().Dispatch("modulexxx", nil, nil) {
		t.Fatal("Dispatch")
	}
	if !client.PathRouterInstance().Dispatch("modulexxxxxxxxxxxxxxx", nil, nil) {
		t.Fatal("Dispatch")
	}
}
