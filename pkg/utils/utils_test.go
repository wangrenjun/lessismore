package utils_test

import (
	"testing"
	"time"

	"github.com/wangrenjun/lessismore/pkg/utils"
)

func fake() bool {
	return true
}

func TestTempfile(t *testing.T) {
	bs := utils.StackTrace(true)
	err := utils.Tempfile("./", time.Now().Format("20060102150405")+"-", bs)
	if err != nil {
		t.Errorf("err: %v", err)
	}

	fn := fake
	t.Logf("function: %s", utils.GetFunctionName(fn))
}
