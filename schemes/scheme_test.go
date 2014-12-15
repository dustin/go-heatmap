package schemes

import (
	"reflect"
	"testing"
)

func TestFromImage(t *testing.T) {
	colors, err := FromImage("nonexistent.png")
	if err == nil {
		t.Errorf("Expected failure on nonexistent flie, got %v", colors)
	}
	colors, err = FromImage("classic.go")
	if err == nil {
		t.Errorf("Expected failure on classic.go, got %v", colors)
	}

	colors, err = FromImage("classic.png")
	if err != nil {
		t.Fatalf("Failed loading classic.png: %v", err)
	}
	if !reflect.DeepEqual(colors, Classic) {
		t.Errorf("Got %v, wanted %v", colors, Classic)
	}
}
