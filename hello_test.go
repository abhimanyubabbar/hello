package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	actual := Hello()
	expect := "Hello, world."

	if expect != actual {
		t.Errorf("Hello(), want:%s, got: %s", expect, actual)
	}
}

func TestProverb(t *testing.T) {
	actual := Proverb()
	expect := "Concurrency is not parallelism."

	if actual != expect {
		t.Errorf("Proverb not working")
	}
}
