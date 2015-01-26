package mogilefs

import (
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	actual := NewClient(map[string]interface{}{
		"Domain": "foo.com::my_namespace",
		"Hosts":  "10.0.0.2:7001",
	})
	expected := "*mogilefs.Client"

	if reflect.TypeOf(actual).String() != expected {
		t.Fatal("failed to create *mogilefs.Client")
	}
}
