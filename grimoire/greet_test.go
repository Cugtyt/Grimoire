package grimoire

import "testing"

func TestGreet(t *testing.T) {
	got := Greet()
	want := "Hello Grimoire!"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}