package di

import (
	"bytes"
	"testing"
)

func Test(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Jatin")

	got := buffer.String()
	want := "Hello, Jatin"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
