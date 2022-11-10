package maps

import (
	"testing"
)

func TestMap(t *testing.T) {
	dictionary := map[string]string{"test": "just a test"}

	got := Search(dictionary, "test")
	want := "just a test"

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
