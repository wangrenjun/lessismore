package utils

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"runtime"
)

func PrettyJson(v interface{}) (marsh string) {
	if b, err := json.MarshalIndent(v, "", "\t"); err == nil {
		marsh = string(b)
	}
	return
}

func StackTrace(all bool) []byte {
	buf := make([]byte, 1<<15)
	for {
		size := runtime.Stack(buf, all)
		if size == len(buf) {
			buf = make([]byte, len(buf)<<1)
			continue
		}
		buf = buf[:size]
		break
	}
	return buf
}

// Remove duplicates in a slice
func RemoveDuplicatesFromSlice(elements []string) []string {
	keys := map[string]struct{}{}
	result := []string{}
	for _, entry := range elements {
		if _, found := keys[entry]; !found {
			keys[entry] = struct{}{}
			result = append(result, entry)
		}
	}
	return result
}

func IsContains(haystack []string, needle string) bool {
	for _, ele := range haystack {
		if ele == needle {
			return true
		}
	}
	return false
}

func Tempfile(path, name string, bytes []byte) (err error) {
	file, err := ioutil.TempFile(path, name)
	if err != nil {
		return
	}
	defer file.Close()
	if _, err = file.Write(bytes); err != nil {
		return
	}
	return
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
