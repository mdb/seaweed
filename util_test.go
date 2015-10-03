package seaweed

import "testing"

func TestConcat(t *testing.T) {
	joined := concat([]string{
		"foo",
		"bar",
	})

	if joined != "foobar" {
		t.Error("concat should properly concatenate strings")
	}
}
