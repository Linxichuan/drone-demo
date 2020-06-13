package go_lang

import "testing"

func TestHello(t *testing.T) {
	if hello() != "hello world" {
		t.Error("Testing error")
	}
}
