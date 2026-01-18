package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	input := "SOmething TO Clean    "
	result := CleanInput(input)
	expected := []string{"something", "to", "clean"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %q, Got: %q", expected, result)
	}

}
