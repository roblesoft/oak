package oak

import (
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	oak := New()
	want := &Oak{
		AppName: "Default",
		server:  &http.ServeMux{},
	}

	if oak != want {
		t.Errorf("Different")
	}
}
