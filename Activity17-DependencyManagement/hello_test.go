package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {

		t.Errorf("Hello() = %q, want %q", got, want)

	} else {
		fmt.Println(got, " vs ", want)
	}

}
