package client_test

import (
	"testing"

	"github.com/wangrenjun/lessismore/internal/client"
)

func TestClientPool(t *testing.T) {
	a := client.NewClient(nil)
	a.Accepted(false)
	a.SetSession(client.NewTemporarySession("a"))
	b := client.NewClient(nil)
	b.Accepted(false)
	b.SetSession(client.NewTemporarySession("b"))
	c := client.NewClient(nil)
	c.Accepted(false)
	c.SetSession(client.NewTemporarySession("c"))
	d := client.NewClient(nil)
	d.Accepted(false)
	d.SetSession(client.NewTemporarySession("d"))

	if client.UnacceptedClientPoolInstance().Num() != 0 {
		t.Fatalf("Num")
	}
	client.UnacceptedClientPoolInstance().Push("1", a)
	client.UnacceptedClientPoolInstance().Push("2", b)
	client.UnacceptedClientPoolInstance().Push("3", c)
	client.UnacceptedClientPoolInstance().Push("4", d)
	if client.UnacceptedClientPoolInstance().Num() != 4 {
		t.Fatalf("Num")
	}

	if !client.UnacceptedClientPoolInstance().Exist("1") {
		t.Fatal("Exist")
	}
	if !client.UnacceptedClientPoolInstance().Exist("2") {
		t.Fatal("Exist")
	}
	if !client.UnacceptedClientPoolInstance().Exist("3") {
		t.Fatal("Exist")
	}
	if !client.UnacceptedClientPoolInstance().Exist("4") {
		t.Fatal("Exist")
	}

	client.UnacceptedClientPoolInstance().Pull("4")
	if client.UnacceptedClientPoolInstance().Exist("4") {
		t.Fatal("Exist")
	}
	if client.UnacceptedClientPoolInstance().Num() != 3 {
		t.Fatalf("Num")
	}

	if client.ClientPoolUnacceptedToAccepted("1", "new1") != a {
		t.Fatal("ClientPoolUnacceptedToAccepted")
	}
	if client.UnacceptedClientPoolInstance().Num() != 2 {
		t.Fatalf("Num")
	}
	if client.ClientPoolUnacceptedToAccepted("1", "new1") != nil {
		t.Fatal("ClientPoolUnacceptedToAccepted")
	}
	if client.AcceptedClientPoolInstance().Num() != 1 {
		t.Fatalf("Num")
	}
	if client.UnacceptedClientPoolInstance().Exist("1") {
		t.Fatal("Exist")
	}
	if !client.AcceptedClientPoolInstance().Exist("new1") {
		t.Fatal("Exist")
	}
	c1, ok := client.AcceptedClientPoolInstance().Pull("new1")
	if !ok || c1 != a {
		t.Fatal("Pull")
	}
	t.Logf("Pull: %s", c1.Session().Id())
	if client.AcceptedClientPoolInstance().Exist("xxxxxxasfdasdfsaaaa") {
		t.Fatal("Exist")
	}
	client.UnacceptedClientPoolInstance().CloseAll()
	if client.UnacceptedClientPoolInstance().Num() != 0 {
		t.Fatalf("Num")
	}

}
