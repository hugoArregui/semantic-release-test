package main

import (
	"testing"
)

func TestSalute(t *testing.T) {
	if Salute() != "Hello world" {
		t.Fatal("error")
	}
}

func TestSalute2(t *testing.T) {
	if Salute() != "Hello world" {
		t.Fatal("error")
	}
}
