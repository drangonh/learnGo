package filelisting

import (
	"errors"
	"fmt"
	"gomodtest/demo/errhandling/filelistingserver/handler"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return handler.UserError("user error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errNotPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnKnown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("error")
}

func noError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Println(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "Error:user error"},
	{errNotFound, 404, "Not Found"},
	{errNotPermission, 403, "Forbidden"},
	{errUnKnown, 500, "Internal Server Error"},
	{noError, 200, ""},
}

/*假的http 测试*/
func TestRrrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := ErrWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8888/list/fib.txt", nil)
		f(response, request)

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

/*用真正的server测试，测试覆盖率更大*/
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := ErrWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)

		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, expectCode int, expectMessage string, t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if body != expectMessage || resp.StatusCode != expectCode {
		t.Errorf("expect:%d,%s,got:%d,%s", expectCode, expectMessage, resp.StatusCode, body)
	}
}
