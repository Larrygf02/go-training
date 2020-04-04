package modules

import "testing"

func TestHello(t *testing.T) {
	want := "Hola Mundo."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
