package health

import (
	"testing"
)

func TestUpdateHealth(t *testing.T) {

	samples := []struct {
		operation byte
		value     int
		player    string
		h1        int
		h2        int
	}{
		{operation: '/', value: 2, player: "p1", h1: 1000, h2: 500},
		{operation: '+', value: 10, player: "p2", h1: 1000, h2: 510},
		{operation: '-', value: 300, player: "p2", h1: 700, h2: 510},
		{operation: '*', value: 3, player: "p1", h1: 1000, h2: 510},
		{operation: '/', value: 5, player: "p2", h1: 200, h2: 510},
		{operation: '+', value: 100, player: "p1", h1: 300, h2: 510},
		{operation: '-', value: 100, player: "p2", h1: 200, h2: 510},
		{operation: '*', value: 2, player: "p2", h1: 200, h2: 1000},
		{operation: '-', value: 100, player: "p2", h1: 100, h2: 1000},
		{operation: '-', value: 500, player: "p1", h1: 100, h2: 500},
		{operation: '/', value: 3, player: "p2", h1: 33, h2: 500},
		{operation: '+', value: 100, player: "p1", h1: 133, h2: 500},
	}

	SetHealth(1000)
	for _, sample := range samples {
		UpdateHealth(sample.operation, sample.value, sample.player)
		if firstHealth != sample.h1 || secondHealth != sample.h2 {
			t.Error("UpdateHealth wasa incorrect, firstHealth:", firstHealth, " secondHealth:", secondHealth, " while operation was ", sample)
		}
	}
}
