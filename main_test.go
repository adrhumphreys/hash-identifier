package main

import "testing"

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestIdentify(t *testing.T) {
	matches, err := identify("5f4dcc3b5aa765d61d8327deb882cf99")

	if err != nil || matches[0].Name != "MD2" {
		t.Fatalf(`identify("5f4dcc3b5aa765d61d8327deb882cf99") = MD2, got %v`, matches)
	}
}
