package example02

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting("Geroge")
	want := "Hello Geroge"
	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}

func TestFarewell(t *testing.T) {
	got := Farewell("Geroge")
	want := "Hello Geroge"
	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}