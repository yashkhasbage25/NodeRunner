package client

import "testing"

func TestGetInfoStr(t *testing.T) {
	cSamples := []struct {
		c      *Client
		answer string
	}{
		{c: &Client{IP: "192.168.0.0", ID: 1, Port: "5678"}, answer: "client: IP: 192.168.0.0\n\t\tPort: 5678\n\t\tID: 1"},
		{c: &Client{IP: "127.0.0.1", ID: 0, Port: "8080"}, answer: "client: IP: 127.0.0.1\n\t\tPort: 8080\n\t\tID: 0"},
	}
	for _, sample := range cSamples {
		got := sample.c.GetInfoStr()
		want := sample.answer
		if got != want {
			t.Error("GetInfoStr was incorrect, got: ", got, "want: ", want)
		}
	}
}
