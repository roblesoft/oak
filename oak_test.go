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

	type testMethodsCases struct {
		description string
		req         *http.Request
		res         *httptest.ResponseRecorder
		want        int
	}

	for _, scenario := range []testMethodsCases{
		{
			description: "Return status code 200",
			req:         httptest.NewRequest(http.MethodGet, "/test", nil),
			res:         httptest.NewRecorder(),
			want:        200,
		},
		{
			description: "Return status code 404",
			req:         httptest.NewRequest(http.MethodPost, "/test", nil),
			res:         httptest.NewRecorder(),
			want:        404,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			oak.GET("/test", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello, World")
			})

			if oak.ServeHTTP(scenario.res, scenario.req); scenario.res.Result().StatusCode != scenario.want {
				t.Errorf("not got status %d", scenario.want)
			}
		})
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
