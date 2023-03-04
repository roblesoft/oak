package oak

import (
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
