package curl

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestServer(serverContent string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, serverContent)
	}))
}

func TestEasyInterface(t *testing.T) {
	ts := setupTestServer("")
	defer ts.Close()

	easy := EasyInit()
	defer easy.Cleanup()

	easy.Setopt(OPT_URL, ts.URL)
	if err := easy.Perform(); err != nil {
		t.Fatal(err)
	}
}

func TestCallbackFunction(t *testing.T) {
	serverContent := "A random string"
	ts := setupTestServer(serverContent)
	defer ts.Close()

	easy := EasyInit()
	defer easy.Cleanup()

	easy.Setopt(OPT_URL, ts.URL)
	easy.Setopt(OPT_WRITEFUNCTION, func(buf []byte, userdata interface{}) bool {
		result := string(buf)
		expected := serverContent + "\n"
		if result != expected {
			t.Errorf("output should be %q and is %q.", expected, result)
		}
		return true
	})
	if err := easy.Perform(); err != nil {
		t.Fatal(err)
	}
}
