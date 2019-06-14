package main

import (
	"testing"
)

func TestListens(t *testing.T) {
	var listens Listens

	listens.Set("1.2.3.4:80")
	listens.Set("1.2.3.4:88")

	if len(listens) != 2 {
		t.Errorf("Length of listens is wrong: got %d, want %d", len(listens), 2)
	} else {
		if listens[0] != "1.2.3.4:80" {
			t.Errorf("listens[0] is wrong: got %s, want %s", listens[0], "1.2.3.4:80")
		}
		if listens[1] != "1.2.3.4:88" {
			t.Errorf("listens[1] is wrong: got %s, want %s", listens[1], "1.2.3.4:88")
		}
	}

	if str := listens.String(); str != "[1.2.3.4:80 1.2.3.4:88]" {
		t.Errorf("listens.String() is wrong: got %s, want %s", str, "[1.2.3.4:80 1.2.3.4:88]")
	}
}
