package codes_test

import (
	"encoding/json"
	"testing"

	"github.com/wangrenjun/lessismore/pkg/codes"
)

func TestCodes(t *testing.T) {
	rc := codes.RC_SERVICE_DOWN
	t.Logf("RC_SERVICE_DOWN: %s", rc)
	if rc.String() != "Service is down" {
		t.Error("Mismatch!!")
	}
	j, err := json.Marshal(rc)
	if err != nil {
		t.Errorf("Marshal: %v", err)
	}
	t.Logf("Json: %s", string(j))

	vv := struct {
		Id   int
		Code codes.ReturnCode
	}{
		1234,
		codes.RC_MSG_TOOL_LARGE,
	}
	j, err = json.Marshal(vv)
	if err != nil {
		t.Errorf("Marshal: %v", err)
	}
	t.Logf("Json: %s", string(j))
}
