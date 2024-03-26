package main

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	if HelloWorld() != expected {
		t.Errorf("HelloWorld() = %v, want %v", HelloWorld(), expected)
	}
}
