package oak

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	oak := New()
	want := &Oak{}

	if reflect.TypeOf(oak) != reflect.TypeOf(want) {
		t.Errorf("Is no returns an oak")
	}
}

func TestGet(t *testing.T) {
	oak := New()

	oak.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		return
	})

	t.Run("Get add route", func(t *testing.T) {
		if len(oak.trees) == 0 {
			t.Errorf("Is no adding routes")
		}

	})
}

func TestServeHTTP(t *testing.T) {
	oak := New()

	oak.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	res := httptest.NewRecorder()

	if oak.ServeHTTP(res, req); res.Result().StatusCode != 200 {
		t.Errorf("not got status 200")
	}
}

func BenchmarkGet(b *testing.B) {
	oak := New()

	oak.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		return
	})
}

func BenchmarkServeHTTP(b *testing.B) {
	oak := New()

	oak.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	res := httptest.NewRecorder()

	oak.ServeHTTP(res, req)
}
