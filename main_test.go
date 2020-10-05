package main

import "testing"

func TestSoma(t *testing.T) {
	soma := soma(5, 5)
	if soma != 10 {
		t.Error("5 + 5 não pode ser igual a ", soma)
	}
}