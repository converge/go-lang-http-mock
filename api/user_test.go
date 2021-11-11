package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func (mock *MockHTTP) Do(_ *http.Request) (*http.Response, error) {
	return mock.response, mock.err
}

type MockHTTP struct {
	response *http.Response
	err error
}

func TestGetRandomUser(t *testing.T) {

	r := httptest.NewRequest(http.MethodGet, "/", nil)

	w := httptest.NewRecorder()

	res := http.Response{
		Body: ioutil.NopCloser(strings.NewReader("hello")),
	}

	mock := MockHTTP{
		response: &res,
	}

	service := Service{
		HTTPClient: &mock,
	}

	service.GetRandomUser(w, r)

	got := w.Result()

	if !reflect.DeepEqual(http.StatusOK, got.StatusCode) {
		t.Errorf("wanted: %d, got: %d", http.StatusOK, got.StatusCode)
	}

}