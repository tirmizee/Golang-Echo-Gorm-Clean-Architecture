package main

import "testing"

func TestHello(t *testing.T) {
	if Hello() != "hello world" {
		t.Error("Expeted 'hello world'")
	}
}
