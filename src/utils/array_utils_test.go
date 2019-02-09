package utils

import "testing"

func TestInArray(t *testing.T) {
	got := InArray("beryllium", []string{"hydrogen", "helium", "lithium"})
	want := false

	if got != want {
		t.Errorf("got '%t' want '%t'", got, want)
	}
}
