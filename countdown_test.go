package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func Countdown(out *bytes.Buffer) {
	fmt.Fprint(out, "3")
}
