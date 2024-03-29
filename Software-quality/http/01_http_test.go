package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"id": 1, "name": "test", "info": "test"}`))
}

func TestMakeHttp(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	want := &Response{
		ID:   1,
		Name: "test",
		Info: "test",
	}

	t.Run("Test MakeHTTPCall", func(t *testing.T) {
		got, err := MakeHTTPCall(server.URL)
		if err != nil {
			t.Errorf("MakeHTTPCall() err = %v; want nil", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MakeHTTPCall() = %v; want %v", got, want)
		}
	})

	fmt.Println(server.URL)
}
