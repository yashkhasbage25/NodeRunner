package client

import "testing"

func TestCompareClientsWithAddr(t *testing.T) {
	got := CompareClientsWithAddr("192.0.0.1", "9000", &Client{IP: "192.1.0.0", Port: "9000"})
	want := false
	if got != want {
		t.Errorf("CompareClientsWithLadder was incorrect, got: %t, want: %t", got, want)
	}
}
