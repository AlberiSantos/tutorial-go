package main

import "testing"

func TestHello(t *testing.T) {
	resultado := Hello("World")
	esperado := "Hello, World!"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
