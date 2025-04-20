package set

import (
	"testing"
)

func TestAddSet(t *testing.T) {
	set := NewSet()

	set.Add("Hello")

	if !set.Has("Hello") {
		t.Errorf("Expected set to contains Hello")
	}
}

func TestRemoveSet(t *testing.T) {
	set := NewSet()

	set.Add("Hello")
	set.Remove("Hello")

	if set.Has("Hello") && set.Size() != 0 {
		t.Errorf("Expected set to be empty")
	}
}

func TestSetSize(t *testing.T) {
	set := NewSet()

	set.Add("1")
	set.Add("2")

	if set.Size() != 2 {
		t.Errorf("Expected set to have a size of 2 but got %d", set.Size())
	}
}

func TestSetHas(t *testing.T) {
	set := NewSet()

	set.Add("Hello")
	set.Add("World")

	if set.Has("Bye") {
		t.Error("Expected that the set would not match")
	}
}

func TestSetDuplicate(t *testing.T) {
	set := NewSet()

	set.Add("Hello")
	set.Add("Hello")

	if set.Size() != 1 {
		t.Errorf("Expected size to be 1 after adding duplicate but got %d", set.Size())
	}
}
