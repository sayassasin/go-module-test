package go_module_test

import (
	"testing"
)

func TestHello(t *testing.T) {
	value := Hello("you")
	if value != "Hello you" {
		t.Errorf("Result is incorrect. got %s, want: %s", value, "Hello you")
	}
}
