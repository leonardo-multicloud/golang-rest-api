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
	err      error
}

func TestGetPrice(t *testing.T) {
	economy := Economy{
		HTTPClient: http.DefaultClient,
	}

	r := httptest.NewRequest(http.MethodGet, "/price", nil)
	w := httptest.NewRecorder()

	economy.GetPrice(w, r)

	got := w.Result()

	if !reflect.DeepEqual(http.StatusOK, got.StatusCode) {
		t.Errorf("want %d, got %d", http.StatusOK, got.StatusCode)
	}
}

func TestGetPriceMock(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/price", nil)
	w := httptest.NewRecorder()

	res := http.Response{
		Body: ioutil.NopCloser(strings.NewReader("Mock")),
	}

	mock := MockHTTP{
		response: &res,
	}

	economy := Economy{
		HTTPClient: &mock,
	}

	economy.GetPrice(w, r)
	got := w.Result()

	if !reflect.DeepEqual(http.StatusOK, got.StatusCode) {
		t.Errorf("want %d, got %d", http.StatusOK, got.StatusCode)
	}
}
