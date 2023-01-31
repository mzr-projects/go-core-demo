package main

import "testing"

func TestHello(t *testing.T) {

	want := "Hello GO!!"

	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want = %q", got, want)
	}

}

func TestHelloQuote(t *testing.T) {

	wantQuote := "Hello, world."

	if got := HelloQuote(); got != wantQuote {
		t.Errorf("HelloQuote() = %q, want = %q", got, wantQuote)
	}
}
