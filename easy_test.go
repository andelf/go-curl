
package curl

import (
	"testing"
)


func TestEasyInterface(t *testing.T) {
	easy := EasyInit()
	defer easy.Cleanup()

	easy.Setopt(OPT_URL, "http://www.baidu.com")
	if err := easy.Perform(); err != nil {
		t.Fatal(err)
	}
}


func TestCallbackFunction(t *testing.T) {
	easy := EasyInit()
	defer easy.Cleanup()

	easy.Setopt(OPT_URL, "http://www.baidu.com")
	easy.Setopt(OPT_WRITEFUNCTION, func (buf []byte, userdata interface{}) bool {
		t.Log("Got: ", string(buf))
		return true
	})
	if err := easy.Perform(); err != nil {
		t.Fatal(err)
	}
}
