package main

import "testing"

func TestEcho(t *testing.T) {
	s := "a"
	r := s + " "
	if r != echo(s) {
		t.Errorf("error")
	}
}
