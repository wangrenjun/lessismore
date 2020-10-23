package pack

import (
	"encoding/json"

	"github.com/wangrenjun/lessismore/pkg/codes"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Packet struct {
}

func Field(packet []byte, name string) (string, bool) {
	r := gjson.Get(string(packet), name)
	return r.String(), r.Exists()
}

func Fields(packet []byte, names ...string) (fields []string) {
	r := gjson.GetMany(string(packet), names...)
	for _, f := range r {
		fields = append(fields, f.String())
	}
	return
}

func PackReply(path string, code codes.ReturnCode, msg interface{}) (rep []byte, err error) {
	if msg != nil {
		rep, err = json.Marshal(msg)
		if err != nil {
			return
		}
	}
	j, err := sjson.Set(string(rep), "Path", path)
	if err != nil {
		return
	}
	j, err = sjson.Set(j, "Code", code)
	if err != nil {
		return
	}
	rep = []byte(j)
	return
}
