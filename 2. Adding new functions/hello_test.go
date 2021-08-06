package hello

import "testing"

func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
}

func TestGlass(t *testing.T) {
	want := "I can eat glass and it doesn't hurt me."
	if got := Glass(); got != want {
		t.Errorf("Glass() = %q, want %q", got, want)
	}
}
