package main

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	if result := HelloWorld(); result != expected {
		t.Errorf("HelloWorld() = %q, want %q", result, expected)
	}
}
