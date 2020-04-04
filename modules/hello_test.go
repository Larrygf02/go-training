package modules

import "testing"

func TestHello(t *testing.T) {
	want := "Hola Mundo."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
}

func TestGoodbye(t *testing.T) {
	want := "Goodbye"
	if got := Goodbye(); got != want {
		t.Errorf("Goodbye() = %q, want %q", got, want)
	}
}
